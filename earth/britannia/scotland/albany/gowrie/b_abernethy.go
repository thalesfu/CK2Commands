package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯内西AbernethyBarony struct {
	feud.BaseBarony
}

var BAbernethy阿伯内西 feud.Barony = &阿伯内西AbernethyBarony{}

func init() {
	f := BAbernethy阿伯内西.(*阿伯内西AbernethyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abernethy",
		TitleName: "阿伯内西",
		TitleCode: "b_abernethy",
	}
}
