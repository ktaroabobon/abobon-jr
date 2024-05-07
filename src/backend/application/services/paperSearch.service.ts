import {Paper} from "../../domain/entities/paper.entity.ts";

export interface PaperSearchService {
  searchPapers(keyword: string): Promise<Paper[]>;
}