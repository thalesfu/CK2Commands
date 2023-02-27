package kemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨武科斯基SavukoskiBarony struct {
	feud.BaseBarony
}

var BSavukoski萨武科斯基 feud.Barony = &萨武科斯基SavukoskiBarony{}

func init() {
    f := BSavukoski萨武科斯基.(*萨武科斯基SavukoskiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "savukoski",
		TitleName: "萨武科斯基",
		TitleCode: "b_savukoski",
	}
}
