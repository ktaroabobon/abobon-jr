import {PaperRepository} from "../../domain/repositories/paper.repository.ts";
import {Paper} from "../../domain/entities/paper.entity.ts";
import {DB} from "https://deno.land/x/sqlite@v3.5.0/mod.ts";

export class PaperRepositoryImpl implements PaperRepository {
  private readonly db: DB;

  constructor(dbName: string) {
    this.db = new DB(dbName);
  }

  async findByDoi(doi: string): Promise<Paper | undefined> {
    const paper = await this.db.query(
      "SELECT * FROM paper WHERE doi = ?",
      [doi]
    );
    return paper[0] as Paper | undefined;
  }

  async save(paper: Paper): Promise<void> {
    await this.db.query(
      "INSERT INTO paper (title, authors, publication_date, publisher, publication_name, doi, naid, url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
      [
        paper.title,
        paper.authors,
        paper.publicationDate,
        paper.publisher,
        paper.publicationName,
        paper.doi,
        paper.naid,
        paper.url,
      ]
    );
  }
}