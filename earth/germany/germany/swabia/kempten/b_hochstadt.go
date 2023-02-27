package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫希施塔特HochstadtBarony struct {
	feud.BaseBarony
}

var BHochstadt赫希施塔特 feud.Barony = &赫希施塔特HochstadtBarony{}

func init() {
    f := BHochstadt赫希施塔特.(*赫希施塔特HochstadtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hochstadt",
		TitleName: "赫希施塔特",
		TitleCode: "b_hochstadt",
	}
}
