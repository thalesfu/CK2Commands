package yatenga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科托拉KotolaBarony struct {
	feud.BaseBarony
}

var BKotola科托拉 feud.Barony = &科托拉KotolaBarony{}

func init() {
	f := BKotola科托拉.(*科托拉KotolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotola",
		TitleName: "科托拉",
		TitleCode: "b_kotola",
	}
}
