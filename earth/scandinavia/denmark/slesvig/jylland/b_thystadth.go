package jylland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐斯泰兹ThystadthBarony struct {
	feud.BaseBarony
}

var BThystadth齐斯泰兹 feud.Barony = &齐斯泰兹ThystadthBarony{}

func init() {
    f := BThystadth齐斯泰兹.(*齐斯泰兹ThystadthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thystadth",
		TitleName: "齐斯泰兹",
		TitleCode: "b_thystadth",
	}
}
