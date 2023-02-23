package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳雷克NarykBarony struct {
	feud.BaseBarony
}

var BNaryk纳雷克 feud.Barony = &纳雷克NarykBarony{}

func init() {
	f := BNaryk纳雷克.(*纳雷克NarykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naryk",
		TitleName: "纳雷克",
		TitleCode: "b_naryk",
	}
}
