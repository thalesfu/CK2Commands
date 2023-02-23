package peresechen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布图切尼ButuceniBarony struct {
	feud.BaseBarony
}

var BButuceni布图切尼 feud.Barony = &布图切尼ButuceniBarony{}

func init() {
	f := BButuceni布图切尼.(*布图切尼ButuceniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "butuceni",
		TitleName: "布图切尼",
		TitleCode: "b_butuceni",
	}
}
