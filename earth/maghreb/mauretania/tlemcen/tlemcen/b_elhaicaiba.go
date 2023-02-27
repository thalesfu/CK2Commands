package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈凯巴ElhaicaibaBarony struct {
	feud.BaseBarony
}

var BElhaicaiba哈凯巴 feud.Barony = &哈凯巴ElhaicaibaBarony{}

func init() {
    f := BElhaicaiba哈凯巴.(*哈凯巴ElhaicaibaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elhaicaiba",
		TitleName: "哈凯巴",
		TitleCode: "b_elhaicaiba",
	}
}
