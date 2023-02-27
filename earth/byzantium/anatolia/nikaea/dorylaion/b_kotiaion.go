package dorylaion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科泰延KotiaionBarony struct {
	feud.BaseBarony
}

var BKotiaion科泰延 feud.Barony = &科泰延KotiaionBarony{}

func init() {
    f := BKotiaion科泰延.(*科泰延KotiaionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotiaion",
		TitleName: "科泰延",
		TitleCode: "b_kotiaion",
	}
}
