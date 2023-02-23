package tyana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉克勒亚AnatoliaherakleaBarony struct {
	feud.BaseBarony
}

var BAnatoliaheraklea赫拉克勒亚 feud.Barony = &赫拉克勒亚AnatoliaherakleaBarony{}

func init() {
	f := BAnatoliaheraklea赫拉克勒亚.(*赫拉克勒亚AnatoliaherakleaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anatoliaheraklea",
		TitleName: "赫拉克勒亚",
		TitleCode: "b_anatoliaheraklea",
	}
}
