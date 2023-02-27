package riga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉斯皮尔斯SalaspilsBarony struct {
	feud.BaseBarony
}

var BSalaspils萨拉斯皮尔斯 feud.Barony = &萨拉斯皮尔斯SalaspilsBarony{}

func init() {
    f := BSalaspils萨拉斯皮尔斯.(*萨拉斯皮尔斯SalaspilsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salaspils",
		TitleName: "萨拉斯皮尔斯",
		TitleCode: "b_salaspils",
	}
}
