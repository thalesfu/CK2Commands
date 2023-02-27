package nurnberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔朗根ErlangenBarony struct {
	feud.BaseBarony
}

var BErlangen埃尔朗根 feud.Barony = &埃尔朗根ErlangenBarony{}

func init() {
    f := BErlangen埃尔朗根.(*埃尔朗根ErlangenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erlangen",
		TitleName: "埃尔朗根",
		TitleCode: "b_erlangen",
	}
}
