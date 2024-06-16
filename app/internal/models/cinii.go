package models

type CiNiiResponse struct {
	Items []struct {
		ID              string   `json:"@id"`
		Title           string   `json:"title"`
		Creators        []string `json:"dc:creator"`
		PublicationDate string   `json:"prism:publicationDate"`
		Publisher       string   `json:"dc:publisher"`
		PublicationName string   `json:"prism:publicationName"`
		Identifiers     []struct {
			Type  string `json:"@type"`
			Value string `json:"@value"`
		} `json:"dc:identifier"`
		Link struct {
			ID string `json:"@id"`
		} `json:"link"`
	}
}
