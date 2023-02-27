package naumadal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫格拉HegraBarony struct {
	feud.BaseBarony
}

var BHegra赫格拉 feud.Barony = &赫格拉HegraBarony{}

func init() {
    f := BHegra赫格拉.(*赫格拉HegraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hegra",
		TitleName: "赫格拉",
		TitleCode: "b_hegra",
	}
}
