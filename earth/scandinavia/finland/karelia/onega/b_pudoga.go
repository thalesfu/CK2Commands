package onega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普多加PudogaBarony struct {
	feud.BaseBarony
}

var BPudoga普多加 feud.Barony = &普多加PudogaBarony{}

func init() {
    f := BPudoga普多加.(*普多加PudogaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pudoga",
		TitleName: "普多加",
		TitleCode: "b_pudoga",
	}
}
