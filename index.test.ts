import { describe, expect, test } from "bun:test";
import { homedir } from "node:os";
import { join } from "node:path";
import { formatSource, isValidSource, pickRandom, SOURCES } from "./index";
import { articles } from "./links";
import {
  addToHistory,
  getReadUrls,
  loadHistory,
  markAsRead,
  markAsUnread,
} from "./storage";

const TEST_CONFIG_DIR = join(homedir(), ".random-go-test");
const _TEST_HISTORY_FILE = join(TEST_CONFIG_DIR, "history.json");

describe("pickRandom", () => {
  test("returns undefined for empty array", () => {
    expect(pickRandom([])).toBeUndefined();
  });

  test("returns the only element for single-element array", () => {
    expect(pickRandom([42])).toBe(42);
  });

  test("returns element from array", () => {
    const arr = [1, 2, 3, 4, 5];
    const result = pickRandom(arr);
    if (result === undefined) throw new Error("Expected result");
    expect(arr).toContain(result);
  });

  test("works with string arrays", () => {
    const arr = ["a", "b", "c"];
    const result = pickRandom(arr);
    if (result === undefined) throw new Error("Expected result");
    expect(arr).toContain(result);
  });

  test("works with object arrays", () => {
    const arr = [{ id: 1 }, { id: 2 }];
    const result = pickRandom(arr);
    if (result === undefined) throw new Error("Expected result");
    expect(arr).toContain(result);
  });
});

describe("formatSource", () => {
  test("formats docs source", () => {
    expect(formatSource("docs")).toBe("Go Docs");
  });

  test("formats tour source", () => {
    expect(formatSource("tour")).toBe("Tour of Go");
  });

  test("formats gobyexample source", () => {
    expect(formatSource("gobyexample")).toBe("Go by Example");
  });

  test("formats pkg source", () => {
    expect(formatSource("pkg")).toBe("Standard Library");
  });

  test("formats blog source", () => {
    expect(formatSource("blog")).toBe("Go Blog");
  });
});

describe("isValidSource", () => {
  test("returns true for valid sources", () => {
    expect(isValidSource("docs")).toBe(true);
    expect(isValidSource("tour")).toBe(true);
    expect(isValidSource("gobyexample")).toBe(true);
    expect(isValidSource("pkg")).toBe(true);
    expect(isValidSource("blog")).toBe(true);
  });

  test("returns false for invalid sources", () => {
    expect(isValidSource("invalid")).toBe(false);
    expect(isValidSource("")).toBe(false);
    expect(isValidSource("DOCS")).toBe(false);
    expect(isValidSource("documentation")).toBe(false);
  });
});

describe("SOURCES constant", () => {
  test("contains all 5 sources", () => {
    expect(SOURCES).toHaveLength(5);
  });

  test("contains expected sources", () => {
    expect(SOURCES).toContain("docs");
    expect(SOURCES).toContain("tour");
    expect(SOURCES).toContain("gobyexample");
    expect(SOURCES).toContain("pkg");
    expect(SOURCES).toContain("blog");
  });
});

describe("articles data", () => {
  test("has articles", () => {
    expect(articles.length).toBeGreaterThan(0);
  });

  test("all articles have required fields", () => {
    for (const article of articles) {
      expect(article.url).toBeDefined();
      expect(typeof article.url).toBe("string");
      expect(article.source).toBeDefined();
      expect(SOURCES).toContain(article.source);
    }
  });

  test("all articles have valid URLs", () => {
    for (const article of articles) {
      expect(article.url.startsWith("http")).toBe(true);
    }
  });

  test("has articles from each source", () => {
    for (const source of SOURCES) {
      const count = articles.filter((a) => a.source === source).length;
      expect(count).toBeGreaterThan(0);
    }
  });
});

describe("storage functions", () => {
  const _testHistoryFile = join(homedir(), ".random-go", "history.json");

  test("loadHistory returns entries array", async () => {
    const history = await loadHistory();
    expect(history).toHaveProperty("entries");
    expect(Array.isArray(history.entries)).toBe(true);
  });

  test("addToHistory adds entry", async () => {
    const testUrl = "https://test.example.com/test-article";
    await addToHistory(testUrl);
    const history = await loadHistory();
    const entry = history.entries.find((e) => e.url === testUrl);
    expect(entry).toBeDefined();
    expect(entry?.isRead).toBe(false);
  });

  test("markAsRead marks entry as read", async () => {
    const testUrl = "https://test.example.com/mark-read-test";
    await addToHistory(testUrl);
    await markAsRead(testUrl);
    const history = await loadHistory();
    const entry = history.entries.find((e) => e.url === testUrl);
    expect(entry?.isRead).toBe(true);
  });

  test("markAsUnread marks entry as unread", async () => {
    const testUrl = "https://test.example.com/unmark-test";
    await addToHistory(testUrl);
    await markAsRead(testUrl);
    await markAsUnread(testUrl);
    const history = await loadHistory();
    const entry = history.entries.find((e) => e.url === testUrl);
    expect(entry?.isRead).toBe(false);
  });

  test("getReadUrls returns Set of read URLs", async () => {
    const readUrls = await getReadUrls();
    expect(readUrls instanceof Set).toBe(true);
  });
});

describe("CLI integration", () => {
  test("help command exits with 0", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "--help"], {
      cwd: import.meta.dir,
    });
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
  });

  test("sources command outputs source list", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "sources"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    await proc.exited;
    expect(output).toContain("docs");
    expect(output).toContain("tour");
    expect(output).toContain("Go Docs");
  });

  test("stats command outputs progress", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "stats"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    await proc.exited;
    expect(output).toContain("Progress");
    expect(output).toContain("Total");
  });

  test("history command works", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "history"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const _output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
  });

  test("random command outputs article", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
    expect(output).toContain("http");
  });

  test("random with source filter works", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "--source=tour"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
    expect(output).toContain("Tour of Go");
  });

  test("invalid source shows error", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "--source=invalid"], {
      cwd: import.meta.dir,
      stderr: "pipe",
    });
    const stderr = await new Response(proc.stderr).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(1);
    expect(stderr).toContain("Invalid source");
  });

  test("mark with index works", async () => {
    await Bun.spawn(["bun", "run", "index.ts"], { cwd: import.meta.dir })
      .exited;

    const proc = Bun.spawn(["bun", "run", "index.ts", "mark", "1"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
    expect(output).toContain("Marked as read");
  });

  test("random with --any flag works", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "--any"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
    expect(output).toContain("http");
  });

  test("explicit random subcommand works", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "random", "-s", "blog"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
    expect(output).toContain("Go Blog");
  });

  test("history with --limit works", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "history", "--limit=2"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
    expect(output).toContain("2 of");
  });

  test("mark with URL works", async () => {
    const testUrl = "https://go.dev/doc/effective_go";
    const proc = Bun.spawn(["bun", "run", "index.ts", "mark", testUrl], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
    expect(output).toContain("Marked as read");
  });

  test("unmark with index works", async () => {
    await Bun.spawn(["bun", "run", "index.ts", "mark", "1"], {
      cwd: import.meta.dir,
    }).exited;

    const proc = Bun.spawn(["bun", "run", "index.ts", "unmark", "1"], {
      cwd: import.meta.dir,
      stdout: "pipe",
    });
    const output = await new Response(proc.stdout).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(0);
    expect(output).toContain("Marked as unread");
  });

  test("mark with invalid index shows error", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "mark", "9999"], {
      cwd: import.meta.dir,
      stderr: "pipe",
    });
    const stderr = await new Response(proc.stderr).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(1);
    expect(stderr).toContain("Invalid index");
  });

  test("unmark with invalid index shows error", async () => {
    const proc = Bun.spawn(["bun", "run", "index.ts", "unmark", "9999"], {
      cwd: import.meta.dir,
      stderr: "pipe",
    });
    const stderr = await new Response(proc.stderr).text();
    const exitCode = await proc.exited;
    expect(exitCode).toBe(1);
    expect(stderr).toContain("Invalid index");
  });
});
