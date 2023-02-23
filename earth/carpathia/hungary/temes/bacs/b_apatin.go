package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿帕廷ApatinBarony struct {
	feud.BaseBarony
}

var BApatin阿帕廷 feud.Barony = &阿帕廷ApatinBarony{}

func init() {
	f := BApatin阿帕廷.(*阿帕廷ApatinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apatin",
		TitleName: "阿帕廷",
		TitleCode: "b_apatin",
	}
}
