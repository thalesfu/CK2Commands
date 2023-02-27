package sieradzka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图雷克TurekBarony struct {
	feud.BaseBarony
}

var BTurek图雷克 feud.Barony = &图雷克TurekBarony{}

func init() {
    f := BTurek图雷克.(*图雷克TurekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turek",
		TitleName: "图雷克",
		TitleCode: "b_turek",
	}
}
