#!/usr/bin/env bun
import { select } from "@inquirer/prompts";
import { defineCommand, runMain } from "citty";
import { articles } from "./links";
import {
  addToHistory,
  getHistoryFilePath,
  getReadUrls,
  loadHistory,
  markAsRead,
  markAsUnread,
} from "./storage";
import type { ArticleSource } from "./types";

export const SOURCES: ArticleSource[] = [
  "docs",
  "tour",
  "gobyexample",
  "pkg",
  "blog",
];

export function pickRandom<T>(arr: T[]): T | undefined {
  if (arr.length === 0) return undefined;
  return arr[Math.floor(Math.random() * arr.length)];
}

export function formatSource(source: ArticleSource): string {
  const names: Record<ArticleSource, string> = {
    docs: "Go Docs",
    tour: "Tour of Go",
    gobyexample: "Go by Example",
    pkg: "Standard Library",
    blog: "Go Blog",
  };
  return names[source];
}

export function isValidSource(source: string): source is ArticleSource {
  return SOURCES.includes(source as ArticleSource);
}

async function commandRandom(
  includeRead: boolean,
  sourceFilter?: string,
): Promise<void> {
  let pool = articles;

  if (sourceFilter) {
    if (!isValidSource(sourceFilter)) {
      console.error(`Invalid source: ${sourceFilter}`);
      console.log(`Available sources: ${SOURCES.join(", ")}`);
      process.exit(1);
    }
    pool = pool.filter((a) => a.source === sourceFilter);
    if (pool.length === 0) {
      console.error(`No articles found for source: ${sourceFilter}`);
      process.exit(1);
    }
  }

  if (!includeRead) {
    const readUrls = await getReadUrls();
    pool = pool.filter((a) => !readUrls.has(a.url));

    if (pool.length === 0) {
      console.log(
        "Congratulations! You have read all articles" +
          (sourceFilter
            ? ` from ${formatSource(sourceFilter as ArticleSource)}`
            : "") +
          "!",
      );
      console.log(
        "Use --any to get a random article from all (including read).",
      );
      return;
    }
  }

  const article = pickRandom(pool);
  if (!article) {
    console.log("No articles available.");
    return;
  }

  await addToHistory(article.url);

  console.log();
  console.log(`[${formatSource(article.source)}]`);
  if (article.title) {
    console.log(`${article.title}`);
  }
  console.log(`${article.url}`);
  console.log();
  console.log('Tip: Use "gorandom mark <url>" to mark as read when done.');
}

async function commandHistory(limit: number): Promise<void> {
  const history = await loadHistory();

  if (history.entries.length === 0) {
    console.log('No history yet. Run "gorandom" to get your first article!');
    return;
  }

  const sorted = [...history.entries].sort(
    (a, b) => new Date(b.viewedAt).getTime() - new Date(a.viewedAt).getTime(),
  );

  const toShow = sorted.slice(0, limit);

  console.log(
    `\nRecent history (${toShow.length} of ${history.entries.length}):\n`,
  );

  for (const entry of toShow) {
    const status = entry.isRead ? "[READ]" : "[    ]";
    const date = new Date(entry.viewedAt).toLocaleDateString();
    const article = articles.find((a) => a.url === entry.url);
    const title = article?.title || "Unknown";
    console.log(`${status} ${date} - ${title}`);
    console.log(`       ${entry.url}`);
  }

  console.log();
}

async function selectFromHistory(filterRead: boolean): Promise<string | null> {
  const history = await loadHistory();
  const sorted = [...history.entries].sort(
    (a, b) => new Date(b.viewedAt).getTime() - new Date(a.viewedAt).getTime(),
  );

  const filtered = filterRead
    ? sorted.filter((e) => e.isRead)
    : sorted.filter((e) => !e.isRead);

  if (filtered.length === 0) {
    console.log(
      filterRead
        ? "No read articles to unmark."
        : "No unread articles to mark.",
    );
    return null;
  }

  const choices = filtered.slice(0, 10).map((entry, i) => {
    const article = articles.find((a) => a.url === entry.url);
    const title = article?.title || entry.url;
    return { name: `${i + 1}. ${title}`, value: entry.url };
  });

  const selected = await select({
    message: filterRead
      ? "Select article to mark as unread:"
      : "Select article to mark as read:",
    choices,
  });

  return selected;
}

async function commandMark(url?: string): Promise<void> {
  let targetUrl: string;

  if (!url) {
    const selected = await selectFromHistory(false);
    if (!selected) return;
    targetUrl = selected;
  } else if (/^\d+$/.test(url)) {
    const history = await loadHistory();
    const sorted = [...history.entries].sort(
      (a, b) => new Date(b.viewedAt).getTime() - new Date(a.viewedAt).getTime(),
    );
    const index = parseInt(url, 10) - 1;
    if (index >= 0 && index < sorted.length) {
      targetUrl = sorted[index]?.url;
    } else {
      console.error(`Invalid index: ${url}. Use 1 for most recent.`);
      process.exit(1);
    }
  } else {
    targetUrl = url;
  }

  await markAsRead(targetUrl);
  const article = articles.find((a) => a.url === targetUrl);
  console.log(`Marked as read: ${article?.title || targetUrl}`);
}

async function commandUnmark(url?: string): Promise<void> {
  let targetUrl: string;

  if (!url) {
    const selected = await selectFromHistory(true);
    if (!selected) return;
    targetUrl = selected;
  } else if (/^\d+$/.test(url)) {
    const history = await loadHistory();
    const sorted = [...history.entries].sort(
      (a, b) => new Date(b.viewedAt).getTime() - new Date(a.viewedAt).getTime(),
    );
    const index = parseInt(url, 10) - 1;
    if (index >= 0 && index < sorted.length) {
      targetUrl = sorted[index]?.url;
    } else {
      console.error(`Invalid index: ${url}. Use 1 for most recent.`);
      process.exit(1);
    }
  } else {
    targetUrl = url;
  }

  const success = await markAsUnread(targetUrl);
  if (success) {
    const article = articles.find((a) => a.url === targetUrl);
    console.log(`Marked as unread: ${article?.title || targetUrl}`);
  } else {
    console.log(`URL not found in history: ${targetUrl}`);
  }
}

async function commandStats(): Promise<void> {
  const readUrls = await getReadUrls();

  console.log("\n📊 Your Go Learning Progress\n");
  console.log("─".repeat(50));

  let totalRead = 0;
  let totalArticles = 0;

  for (const source of SOURCES) {
    const sourceArticles = articles.filter((a) => a.source === source);
    const sourceRead = sourceArticles.filter((a) => readUrls.has(a.url)).length;
    totalRead += sourceRead;
    totalArticles += sourceArticles.length;

    const percent =
      sourceArticles.length > 0
        ? Math.round((sourceRead / sourceArticles.length) * 100)
        : 0;
    const bar =
      "█".repeat(Math.floor(percent / 5)) +
      "░".repeat(20 - Math.floor(percent / 5));

    console.log(
      `${formatSource(source).padEnd(18)} ${bar} ${sourceRead.toString().padStart(3)}/${sourceArticles.length.toString().padStart(3)} (${percent}%)`,
    );
  }

  console.log("─".repeat(50));
  const totalPercent =
    totalArticles > 0 ? Math.round((totalRead / totalArticles) * 100) : 0;
  console.log(
    `${"Total".padEnd(18)} ${" ".repeat(20)} ${totalRead.toString().padStart(3)}/${totalArticles.toString().padStart(3)} (${totalPercent}%)`,
  );
  console.log();
  console.log(`History stored at: ${getHistoryFilePath()}`);
  console.log();
}

function commandSources(): void {
  console.log("\nAvailable sources:\n");
  for (const source of SOURCES) {
    const count = articles.filter((a) => a.source === source).length;
    console.log(
      `  ${source.padEnd(15)} - ${formatSource(source)} (${count} articles)`,
    );
  }
  console.log();
}

const historyCommand = defineCommand({
  meta: {
    name: "history",
    description: "Show viewing history",
  },
  args: {
    limit: {
      type: "string",
      description: "Number of entries to show",
      default: "10",
    },
  },
  run: async ({ args }) => {
    const limit = parseInt(args.limit, 10) || 10;
    await commandHistory(limit);
  },
});

const markCommand = defineCommand({
  meta: {
    name: "mark",
    description: "Mark article as read (interactive if no arg)",
  },
  args: {
    target: {
      type: "positional",
      description: "URL or history index (1 = most recent)",
      required: false,
    },
  },
  run: async ({ args }) => {
    await commandMark(args.target);
  },
});

const unmarkCommand = defineCommand({
  meta: {
    name: "unmark",
    description: "Mark article as unread (interactive if no arg)",
  },
  args: {
    target: {
      type: "positional",
      description: "URL or history index (1 = most recent)",
      required: false,
    },
  },
  run: async ({ args }) => {
    await commandUnmark(args.target);
  },
});

const statsCommand = defineCommand({
  meta: {
    name: "stats",
    description: "Show read/unread counts per source",
  },
  run: async () => {
    await commandStats();
  },
});

const sourcesCommand = defineCommand({
  meta: {
    name: "sources",
    description: "List available sources",
  },
  run: () => {
    commandSources();
  },
});

const randomCommand = defineCommand({
  meta: {
    name: "random",
    description: "Get a random article (default command)",
  },
  args: {
    any: {
      type: "boolean",
      alias: "a",
      description: "Include read articles",
      default: false,
    },
    source: {
      type: "string",
      alias: "s",
      description: "Filter by source (docs, tour, gobyexample, pkg, blog)",
    },
  },
  run: async ({ args }) => {
    await commandRandom(args.any, args.source);
  },
});

const SUBCOMMAND_NAMES = [
  "random",
  "history",
  "mark",
  "unmark",
  "stats",
  "sources",
];

const main = defineCommand({
  meta: {
    name: "gorandom",
    version: "1.0.0",
    description: "Get random Go learning articles",
  },
  args: {
    any: {
      type: "boolean",
      alias: "a",
      description: "Include read articles",
      default: false,
    },
    source: {
      type: "string",
      alias: "s",
      description: 'Filter by source (use --source=VALUE or "random -s VALUE")',
    },
  },
  subCommands: {
    random: randomCommand,
    history: historyCommand,
    mark: markCommand,
    unmark: unmarkCommand,
    stats: statsCommand,
    sources: sourcesCommand,
  },
  run: async ({ args, rawArgs }) => {
    const hasSubcommand = rawArgs.some((arg) => SUBCOMMAND_NAMES.includes(arg));
    if (!hasSubcommand) {
      await commandRandom(args.any, args.source);
    }
  },
});

runMain(main);
