import {RssItemRepository} from "../../domain/repositories/rssItem.repository.ts";
import {RssItem} from "../../domain/entities/rssItem.entity.ts";

export class FetchLatestRssItemsUseCase {
  constructor(private readonly rssItemRepository: RssItemRepository) {
  }

  execute(url: string, limit: number): Promise<RssItem[]> {
    return this.rssItemRepository.fetchLatestItems(url, limit);
  }
}