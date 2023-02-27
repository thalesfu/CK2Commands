package jaunpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豆利揭DohrighatBarony struct {
	feud.BaseBarony
}

var BDohrighat豆利揭 feud.Barony = &豆利揭DohrighatBarony{}

func init() {
    f := BDohrighat豆利揭.(*豆利揭DohrighatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dohrighat",
		TitleName: "豆利揭",
		TitleCode: "b_dohrighat",
	}
}
