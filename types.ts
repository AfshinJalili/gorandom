export type ArticleSource = 'docs' | 'tour' | 'gobyexample' | 'pkg' | 'blog';

export interface Article {
  url: string;
  source: ArticleSource;
  title?: string;
}

export interface HistoryEntry {
  url: string;
  viewedAt: string;
  isRead: boolean; // marked as "completed"
}

export interface HistoryData {
  entries: HistoryEntry[];
}
