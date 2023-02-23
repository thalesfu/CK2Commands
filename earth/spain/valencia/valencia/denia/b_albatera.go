package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔瓦特拉AlbateraBarony struct {
	feud.BaseBarony
}

var BAlbatera阿尔瓦特拉 feud.Barony = &阿尔瓦特拉AlbateraBarony{}

func init() {
	f := BAlbatera阿尔瓦特拉.(*阿尔瓦特拉AlbateraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albatera",
		TitleName: "阿尔瓦特拉",
		TitleCode: "b_albatera",
	}
}
