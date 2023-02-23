package chaldea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯拉苏斯KerasousBarony struct {
	feud.BaseBarony
}

var BKerasous凯拉苏斯 feud.Barony = &凯拉苏斯KerasousBarony{}

func init() {
	f := BKerasous凯拉苏斯.(*凯拉苏斯KerasousBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerasous",
		TitleName: "凯拉苏斯",
		TitleCode: "b_kerasous",
	}
}
