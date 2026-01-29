#!/usr/bin/env bun
import { articles } from './links';
import {
  loadHistory,
  addToHistory,
  markAsRead,
  markAsUnread,
  getReadUrls,
  getHistoryFilePath,
} from './storage';
import type { Article, ArticleSource } from './types';

const SOURCES: ArticleSource[] = ['docs', 'tour', 'gobyexample', 'pkg', 'blog'];

function printHelp(): void {
  console.log(`
gorandom- Get random Go learning articles

USAGE:
  gorandom                    Get a random unread article
  gorandom--any               Get any random article (including read)
  gorandom--source <source>   Filter by source (docs, tour, gobyexample, pkg, blog)
  gorandomhistory             Show viewing history
  gorandomhistory --limit N   Show last N history entries (default: 10)
  gorandommark <url>          Mark article as read
  gorandomunmark <url>        Mark article as unread
  gorandomstats               Show read/unread counts per source
  gorandomsources             List available sources
  gorandomhelp                Show this help message

EXAMPLES:
  gorandom                    # Get a random unread article
  gorandom--source tour       # Get a random Tour of Go page
  gorandom--any --source blog # Get any random blog post
  gorandommark https://...    # Mark a URL as read
  gorandomstats               # See your progress
`);
}

function pickRandom<T>(arr: T[]): T | undefined {
  if (arr.length === 0) return undefined;
  return arr[Math.floor(Math.random() * arr.length)];
}

function formatSource(source: ArticleSource): string {
  const names: Record<ArticleSource, string> = {
    docs: 'Go Docs',
    tour: 'Tour of Go',
    gobyexample: 'Go by Example',
    pkg: 'Standard Library',
    blog: 'Go Blog',
  };
  return names[source];
}

async function commandRandom(includeRead: boolean, sourceFilter?: ArticleSource): Promise<void> {
  let pool = articles;
  
  // Filter by source if specified
  if (sourceFilter) {
    pool = pool.filter(a => a.source === sourceFilter);
    if (pool.length === 0) {
      console.error(`No articles found for source: ${sourceFilter}`);
      console.log(`Available sources: ${SOURCES.join(', ')}`);
      process.exit(1);
    }
  }
  
  // Filter out read articles unless --any is specified
  if (!includeRead) {
    const readUrls = await getReadUrls();
    pool = pool.filter(a => !readUrls.has(a.url));
    
    if (pool.length === 0) {
      console.log('Congratulations! You have read all articles' + (sourceFilter ? ` from ${formatSource(sourceFilter)}` : '') + '!');
      console.log('Use --any to get a random article from all (including read).');
      return;
    }
  }
  
  const article = pickRandom(pool);
  if (!article) {
    console.log('No articles available.');
    return;
  }
  
  // Add to history
  await addToHistory(article.url);
  
  // Print the article
  console.log();
  console.log(`[${formatSource(article.source)}]`);
  if (article.title) {
    console.log(`${article.title}`);
  }
  console.log(`${article.url}`);
  console.log();
  console.log('Tip: Use "gorandommark <url>" to mark as read when done.');
}

async function commandHistory(limit: number): Promise<void> {
  const history = await loadHistory();
  
  if (history.entries.length === 0) {
    console.log('No history yet. Run "random-go" to get your first article!');
    return;
  }
  
  // Sort by most recent first
  const sorted = [...history.entries].sort(
    (a, b) => new Date(b.viewedAt).getTime() - new Date(a.viewedAt).getTime()
  );
  
  const toShow = sorted.slice(0, limit);
  
  console.log(`\nRecent history (${toShow.length} of ${history.entries.length}):\n`);
  
  for (const entry of toShow) {
    const status = entry.isRead ? '[READ]' : '[    ]';
    const date = new Date(entry.viewedAt).toLocaleDateString();
    const article = articles.find(a => a.url === entry.url);
    const title = article?.title || 'Unknown';
    console.log(`${status} ${date} - ${title}`);
    console.log(`       ${entry.url}`);
  }
  
  console.log();
}

async function commandMark(url: string): Promise<void> {
  // Try to find article by URL or partial match
  let targetUrl = url;
  
  // If it's a number, treat it as "mark last" shortcut
  if (/^\d+$/.test(url)) {
    const history = await loadHistory();
    const sorted = [...history.entries].sort(
      (a, b) => new Date(b.viewedAt).getTime() - new Date(a.viewedAt).getTime()
    );
    const index = parseInt(url, 10) - 1;
    if (index >= 0 && index < sorted.length) {
      targetUrl = sorted[index].url;
    } else {
      console.error(`Invalid index: ${url}. Use 1 for most recent.`);
      process.exit(1);
    }
  }
  
  await markAsRead(targetUrl);
  const article = articles.find(a => a.url === targetUrl);
  console.log(`Marked as read: ${article?.title || targetUrl}`);
}

async function commandUnmark(url: string): Promise<void> {
  let targetUrl = url;
  
  if (/^\d+$/.test(url)) {
    const history = await loadHistory();
    const sorted = [...history.entries].sort(
      (a, b) => new Date(b.viewedAt).getTime() - new Date(a.viewedAt).getTime()
    );
    const index = parseInt(url, 10) - 1;
    if (index >= 0 && index < sorted.length) {
      targetUrl = sorted[index].url;
    } else {
      console.error(`Invalid index: ${url}. Use 1 for most recent.`);
      process.exit(1);
    }
  }
  
  const success = await markAsUnread(targetUrl);
  if (success) {
    const article = articles.find(a => a.url === targetUrl);
    console.log(`Marked as unread: ${article?.title || targetUrl}`);
  } else {
    console.log(`URL not found in history: ${targetUrl}`);
  }
}

async function commandStats(): Promise<void> {
  const readUrls = await getReadUrls();
  
  console.log('\n📊 Your Go Learning Progress\n');
  console.log('─'.repeat(50));
  
  let totalRead = 0;
  let totalArticles = 0;
  
  for (const source of SOURCES) {
    const sourceArticles = articles.filter(a => a.source === source);
    const sourceRead = sourceArticles.filter(a => readUrls.has(a.url)).length;
    totalRead += sourceRead;
    totalArticles += sourceArticles.length;
    
    const percent = sourceArticles.length > 0 
      ? Math.round((sourceRead / sourceArticles.length) * 100) 
      : 0;
    const bar = '█'.repeat(Math.floor(percent / 5)) + '░'.repeat(20 - Math.floor(percent / 5));
    
    console.log(`${formatSource(source).padEnd(18)} ${bar} ${sourceRead.toString().padStart(3)}/${sourceArticles.length.toString().padStart(3)} (${percent}%)`);
  }
  
  console.log('─'.repeat(50));
  const totalPercent = totalArticles > 0 
    ? Math.round((totalRead / totalArticles) * 100) 
    : 0;
  console.log(`${'Total'.padEnd(18)} ${' '.repeat(20)} ${totalRead.toString().padStart(3)}/${totalArticles.toString().padStart(3)} (${totalPercent}%)`);
  console.log();
  console.log(`History stored at: ${getHistoryFilePath()}`);
  console.log();
}

function commandSources(): void {
  console.log('\nAvailable sources:\n');
  for (const source of SOURCES) {
    const count = articles.filter(a => a.source === source).length;
    console.log(`  ${source.padEnd(15)} - ${formatSource(source)} (${count} articles)`);
  }
  console.log();
}

async function main(): Promise<void> {
  const args = process.argv.slice(2);
  
  if (args.length === 0) {
    // Default: get random unread article
    await commandRandom(false);
    return;
  }
  
  const command = args[0];
  
  switch (command) {
    case 'help':
    case '--help':
    case '-h':
      printHelp();
      break;
      
    case 'history':
      let limit = 10;
      const limitIndex = args.indexOf('--limit');
      if (limitIndex !== -1 && args[limitIndex + 1]) {
        limit = parseInt(args[limitIndex + 1], 10) || 10;
      }
      await commandHistory(limit);
      break;
      
    case 'mark':
      if (!args[1]) {
        console.error('Usage: gorandommark <url or index>');
        console.log('Use index 1 for most recent, 2 for second most recent, etc.');
        process.exit(1);
      }
      await commandMark(args[1]);
      break;
      
    case 'unmark':
      if (!args[1]) {
        console.error('Usage: gorandomunmark <url or index>');
        process.exit(1);
      }
      await commandUnmark(args[1]);
      break;
      
    case 'stats':
      await commandStats();
      break;
      
    case 'sources':
      commandSources();
      break;
      
    default:
      // Check for flags
      let includeRead = false;
      let sourceFilter: ArticleSource | undefined;
      
      for (let i = 0; i < args.length; i++) {
        if (args[i] === '--any' || args[i] === '-a') {
          includeRead = true;
        } else if ((args[i] === '--source' || args[i] === '-s') && args[i + 1]) {
          const source = args[i + 1] as ArticleSource;
          if (!SOURCES.includes(source)) {
            console.error(`Invalid source: ${source}`);
            console.log(`Available sources: ${SOURCES.join(', ')}`);
            process.exit(1);
          }
          sourceFilter = source;
          i++; // Skip next arg
        }
      }
      
      await commandRandom(includeRead, sourceFilter);
  }
}

main().catch(console.error);
