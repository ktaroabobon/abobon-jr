import { Paper } from "../entities/paper.entity.ts";

export interface PaperRepository {
  findByDoi(doi: string): Promise<Paper | undefined>;
  save(paper: Paper): Promise<void>;
}