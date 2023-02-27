package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加兹温QazwinBarony struct {
	feud.BaseBarony
}

var BQazwin加兹温 feud.Barony = &加兹温QazwinBarony{}

func init() {
    f := BQazwin加兹温.(*加兹温QazwinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qazwin",
		TitleName: "加兹温",
		TitleCode: "b_qazwin",
	}
}
