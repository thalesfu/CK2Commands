package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯法拉延IsfarayenBarony struct {
	feud.BaseBarony
}

var BIsfarayen埃斯法拉延 feud.Barony = &埃斯法拉延IsfarayenBarony{}

func init() {
	f := BIsfarayen埃斯法拉延.(*埃斯法拉延IsfarayenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isfarayen",
		TitleName: "埃斯法拉延",
		TitleCode: "b_isfarayen",
	}
}
