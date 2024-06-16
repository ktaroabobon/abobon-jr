package models

import "time"

// Paper 構造体は学術論文のデータを表します。
// ID, Title, DOI, URL は必須フィールドです。
type Paper struct {
	ID              uint       // 論文の一意識別子、必須
	Title           string     // 論文のタイトル、必須
	Authors         *string    // 論文の著者、任意
	PublicationDate *time.Time // 論文の公開日、任意
	Publisher       *string    // 出版社、任意
	PublicationName *string    // 出版物名、任意
	DOI             string     // デジタルオブジェクト識別子、必須
	NAID            *string    // 国立情報学研究所の識別子、任意
	URL             string     // 論文のURL、必須
}
