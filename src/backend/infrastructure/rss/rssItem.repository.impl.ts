import {RssItemRepository} from "../../domain/repositories/rssItem.repository.ts";
import {RssItem} from "../../domain/entities/rssItem.entity.ts";
import {parseFeed, DublinCore} from "../../../../deps.ts";

export class RssItemRepositoryImpl implements RssItemRepository {
  async fetchLatestItems(url: string, limit: number): Promise<RssItem[]> {
    try {
      const response = await fetch(url);
      const xml = await response.text();
      const feed = await parseFeed(xml);
      return feed.entries.slice(0, limit).map((entry) => ({
        title: entry.title ?? entry[DublinCore.Title] ?? "",
        link: entry.links?.at(0) ?? entry[DublinCore.URI] ?? "",
      }));
    } catch (error) {
      console.error("Failed to fetch RSS feed:", error);
      return [];
    }
  }
}