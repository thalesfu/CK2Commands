package vasyugan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 万普戈尔VampugolBarony struct {
	feud.BaseBarony
}

var BVampugol万普戈尔 feud.Barony = &万普戈尔VampugolBarony{}

func init() {
    f := BVampugol万普戈尔.(*万普戈尔VampugolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vampugol",
		TitleName: "万普戈尔",
		TitleCode: "b_vampugol",
	}
}
