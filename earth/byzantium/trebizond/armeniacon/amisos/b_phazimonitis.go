package amisos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法齐莫尼蒂斯PhazimonitisBarony struct {
	feud.BaseBarony
}

var BPhazimonitis法齐莫尼蒂斯 feud.Barony = &法齐莫尼蒂斯PhazimonitisBarony{}

func init() {
	f := BPhazimonitis法齐莫尼蒂斯.(*法齐莫尼蒂斯PhazimonitisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phazimonitis",
		TitleName: "法齐莫尼蒂斯",
		TitleCode: "b_phazimonitis",
	}
}
