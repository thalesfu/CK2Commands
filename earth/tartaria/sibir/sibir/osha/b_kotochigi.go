package osha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科托奇吉KotochigiBarony struct {
	feud.BaseBarony
}

var BKotochigi科托奇吉 feud.Barony = &科托奇吉KotochigiBarony{}

func init() {
	f := BKotochigi科托奇吉.(*科托奇吉KotochigiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotochigi",
		TitleName: "科托奇吉",
		TitleCode: "b_kotochigi",
	}
}
