import {PaperSearchService} from "../../application/services/paperSearch.service.ts";
import {Paper} from "../../domain/entities/paper.entity.ts";

interface CiNiiIdentifier {
  "@type": string;
  "@value": string;
}

export class PaperSearchServiceImpl implements PaperSearchService {
  async searchPapers(keyword: string): Promise<Paper[]> {
    const url = new URL("https://cir.nii.ac.jp/opensearch/articles");
    url.searchParams.append("q", encodeURIComponent(keyword));
    url.searchParams.append("format", "json");
    url.searchParams.append("count", "5");
    url.searchParams.append("lang", "ja");
    url.searchParams.append("appid", this.appId);

    const response = await fetch(url.toString());
    const data = await response.json();

    return data.items.map((item: {
      "@id": string;
      title?: string;
      "dc:creator"?: string[];
      "prism:publicationDate"?: string;
      "dc:publisher"?: string;
      "prism:publicationName"?: string;
      "dc:identifier"?: CiNiiIdentifier[];
      link?: { "@id": string };
    }) => ({
      id: item["@id"],
      title: item.title ?? "",
      authors: item["dc:creator"]?.join(", ") ?? "",
      publicationDate: item["prism:publicationDate"] ?? "",
      publisher: item["dc:publisher"] ?? "",
      publicationName: item["prism:publicationName"] ?? "",
      doi: item["dc:identifier"]?.find((id) => id["@type"] === "cir:DOI")?.["@value"] ?? "",
      naid: item["dc:identifier"]?.find((id) => id["@type"] === "cir:NAID")?.["@value"] ?? "",
      url: item.link?.["@id"] ?? "",
    }));
  }
}