package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯米亚多沃SmyadovoBarony struct {
	feud.BaseBarony
}

var BSmyadovo斯米亚多沃 feud.Barony = &斯米亚多沃SmyadovoBarony{}

func init() {
	f := BSmyadovo斯米亚多沃.(*斯米亚多沃SmyadovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "smyadovo",
		TitleName: "斯米亚多沃",
		TitleCode: "b_smyadovo",
	}
}
