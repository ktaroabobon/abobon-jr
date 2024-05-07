import {RssItem} from "../entities/rssItem.entity.ts";

export interface RssItemRepository {
  fetchLatestItems(url: string, limit: number): Promise<RssItem[]>;
}