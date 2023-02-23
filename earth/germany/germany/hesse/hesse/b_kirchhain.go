package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基希海恩KirchhainBarony struct {
	feud.BaseBarony
}

var BKirchhain基希海恩 feud.Barony = &基希海恩KirchhainBarony{}

func init() {
	f := BKirchhain基希海恩.(*基希海恩KirchhainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirchhain",
		TitleName: "基希海恩",
		TitleCode: "b_kirchhain",
	}
}
