package dyfed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳伯斯NarberthBarony struct {
	feud.BaseBarony
}

var BNarberth纳伯斯 feud.Barony = &纳伯斯NarberthBarony{}

func init() {
	f := BNarberth纳伯斯.(*纳伯斯NarberthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narberth",
		TitleName: "纳伯斯",
		TitleCode: "b_narberth",
	}
}
