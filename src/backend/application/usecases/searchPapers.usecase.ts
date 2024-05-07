import {PaperSearchService} from "../services/paperSearch.service.ts";
import {PaperRepository} from "../../domain/repositories/paper.repository.ts";
import {Paper} from "../../domain/entities/paper.entity.ts";

export class SearchPapersUseCase {
  constructor(
    private readonly paperSearchService: PaperSearchService,
    private readonly paperRepository: PaperRepository
  ) {
  }

  async execute(keyword: string): Promise<Paper[]> {
    const papers = await this.paperSearchService.searchPapers(keyword);
    const newPapers = await Promise.all(
      papers.map(async (paper) => {
        const existingPaper = await this.paperRepository.findByDoi(paper.doi);
        if (!existingPaper) {
          await this.paperRepository.save(paper);
          return paper;
        }
        return null;
      })
    );
    return newPapers.filter((paper): paper is Paper => paper !== null);
  }
}