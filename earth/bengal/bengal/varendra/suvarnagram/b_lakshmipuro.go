package suvarnagram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗迦施弥补罗LakshmipuroBarony struct {
	feud.BaseBarony
}

var BLakshmipuro罗迦施弥补罗 feud.Barony = &罗迦施弥补罗LakshmipuroBarony{}

func init() {
	f := BLakshmipuro罗迦施弥补罗.(*罗迦施弥补罗LakshmipuroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lakshmipuro",
		TitleName: "罗迦施弥补罗",
		TitleCode: "b_lakshmipuro",
	}
}
