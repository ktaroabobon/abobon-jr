import { Paper } from "../../domain/entities/paper.entity.ts";
import { PaperRepository } from "../../domain/repositories/paper.repository.ts";
import { Database, MySQLConnector } from "../../../deps.ts";

export class PaperRepositoryImpl implements PaperRepository {
  private db: Database;

  constructor(private readonly connectionString: string) {
    this.db = new Database(new MySQLConnector({
      uri: connectionString,
    }));
    this.db.link([Paper]);
  }

  async connect() {
    await this.db.sync();
  }

  async disconnect() {
    await this.db.close();
  }

  async findByDoi(doi: string): Promise<Paper | undefined> {
    return await Paper.where({ doi }).first();
  }

  async save(paper: Paper): Promise<void> {
    await Paper.create({
      title: paper.title,
      authors: paper.authors,
      publicationDate: paper.publicationDate,
      publisher: paper.publisher,
      publicationName: paper.publicationName,
      doi: paper.doi,
      naid: paper.naid,
      url: paper.url,
    });
  }
}
