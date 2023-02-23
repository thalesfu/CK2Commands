package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普雷韦扎PrevezaBarony struct {
	feud.BaseBarony
}

var BPreveza普雷韦扎 feud.Barony = &普雷韦扎PrevezaBarony{}

func init() {
	f := BPreveza普雷韦扎.(*普雷韦扎PrevezaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "preveza",
		TitleName: "普雷韦扎",
		TitleCode: "b_preveza",
	}
}
