package mstislavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗加乔夫RahacouBarony struct {
	feud.BaseBarony
}

var BRahacou罗加乔夫 feud.Barony = &罗加乔夫RahacouBarony{}

func init() {
	f := BRahacou罗加乔夫.(*罗加乔夫RahacouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rahacou",
		TitleName: "罗加乔夫",
		TitleCode: "b_rahacou",
	}
}
