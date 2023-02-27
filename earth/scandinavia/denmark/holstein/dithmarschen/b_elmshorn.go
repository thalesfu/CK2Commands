package dithmarschen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔姆斯霍恩ElmshornBarony struct {
	feud.BaseBarony
}

var BElmshorn埃尔姆斯霍恩 feud.Barony = &埃尔姆斯霍恩ElmshornBarony{}

func init() {
    f := BElmshorn埃尔姆斯霍恩.(*埃尔姆斯霍恩ElmshornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elmshorn",
		TitleName: "埃尔姆斯霍恩",
		TitleCode: "b_elmshorn",
	}
}
