package acalapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉姆KalamBarony struct {
	feud.BaseBarony
}

var BKalam加拉姆 feud.Barony = &加拉姆KalamBarony{}

func init() {
	f := BKalam加拉姆.(*加拉姆KalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalam",
		TitleName: "加拉姆",
		TitleCode: "b_kalam",
	}
}
