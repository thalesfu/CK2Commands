package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特维StewrBarony struct {
	feud.BaseBarony
}

var BStewr斯特维 feud.Barony = &斯特维StewrBarony{}

func init() {
    f := BStewr斯特维.(*斯特维StewrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stewr",
		TitleName: "斯特维",
		TitleCode: "b_stewr",
	}
}
