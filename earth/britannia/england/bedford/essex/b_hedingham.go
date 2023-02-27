package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫丁厄姆HedinghamBarony struct {
	feud.BaseBarony
}

var BHedingham赫丁厄姆 feud.Barony = &赫丁厄姆HedinghamBarony{}

func init() {
    f := BHedingham赫丁厄姆.(*赫丁厄姆HedinghamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hedingham",
		TitleName: "赫丁厄姆",
		TitleCode: "b_hedingham",
	}
}
