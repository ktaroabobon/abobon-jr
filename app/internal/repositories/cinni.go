package repositories

import (
	"app/internal/models"
	"app/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CiNiiRepository struct {
	CiNiiAppId string
	Logger     *utils.Logger
}

func NewCiNiiRepository(appId string, logger *utils.Logger) *CiNiiRepository {
	return &CiNiiRepository{
		CiNiiAppId: appId,
		Logger:     logger,
	}
}

func (r *CiNiiRepository) FetchThesis(keyword string) (*models.CiNiiResponse, error) {
	// リクエストパラメータの設定
	baseURL := "https://cir.nii.ac.jp/opensearch/articles"
	params := url.Values{}
	params.Add("q", keyword)
	params.Add("format", "json")
	params.Add("count", "5")
	params.Add("lang", "ja")
	params.Add("start", "1")
	params.Add("appid", r.CiNiiAppId)

	requestURL := baseURL + "?" + params.Encode()
	r.Logger.InfoLogger.Printf("Request URL: %s", requestURL)

	// リクエストURLの生成
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return &models.CiNiiResponse{}, fmt.Errorf("CiNii APIへのリクエスト中にエラーが発生しました: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &models.CiNiiResponse{}, fmt.Errorf("CiNii APIへのリクエスト中にエラーが発生しました: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &models.CiNiiResponse{}, fmt.Errorf("CiNii APIからのレスポンスが正常ではありません: %v", resp.Status)
	}

	var data models.CiNiiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return &models.CiNiiResponse{}, fmt.Errorf("JSONデコードエラー: %v", err)
	}

	return &data, nil
}
