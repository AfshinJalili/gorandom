import { existsSync } from "node:fs";
import { mkdir, readFile, writeFile } from "node:fs/promises";
import { homedir } from "node:os";
import { join } from "node:path";
import type { HistoryData } from "./types";

const CONFIG_DIR = join(homedir(), ".random-go");
const HISTORY_FILE = join(CONFIG_DIR, "history.json");

async function ensureConfigDir(): Promise<void> {
  if (!existsSync(CONFIG_DIR)) {
    await mkdir(CONFIG_DIR, { recursive: true });
  }
}

export async function loadHistory(): Promise<HistoryData> {
  await ensureConfigDir();

  if (!existsSync(HISTORY_FILE)) {
    return { entries: [] };
  }

  try {
    const data = await readFile(HISTORY_FILE, "utf-8");
    return JSON.parse(data) as HistoryData;
  } catch {
    return { entries: [] };
  }
}

export async function saveHistory(history: HistoryData): Promise<void> {
  await ensureConfigDir();
  await writeFile(HISTORY_FILE, JSON.stringify(history, null, 2), "utf-8");
}

export async function addToHistory(url: string): Promise<void> {
  const history = await loadHistory();

  // Check if already in history
  const existing = history.entries.find((e) => e.url === url);
  if (existing) {
    // Update viewedAt time
    existing.viewedAt = new Date().toISOString();
  } else {
    // Add new entry
    history.entries.push({
      url,
      viewedAt: new Date().toISOString(),
      isRead: false,
    });
  }

  await saveHistory(history);
}

export async function markAsRead(url: string): Promise<boolean> {
  const history = await loadHistory();

  const entry = history.entries.find((e) => e.url === url);
  if (entry) {
    entry.isRead = true;
    await saveHistory(history);
    return true;
  }

  // If not in history, add it as read
  history.entries.push({
    url,
    viewedAt: new Date().toISOString(),
    isRead: true,
  });
  await saveHistory(history);
  return true;
}

export async function markAsUnread(url: string): Promise<boolean> {
  const history = await loadHistory();

  const entry = history.entries.find((e) => e.url === url);
  if (entry) {
    entry.isRead = false;
    await saveHistory(history);
    return true;
  }

  return false;
}

export async function getReadUrls(): Promise<Set<string>> {
  const history = await loadHistory();
  return new Set(history.entries.filter((e) => e.isRead).map((e) => e.url));
}

export async function getViewedUrls(): Promise<Set<string>> {
  const history = await loadHistory();
  return new Set(history.entries.map((e) => e.url));
}

export function getHistoryFilePath(): string {
  return HISTORY_FILE;
}
