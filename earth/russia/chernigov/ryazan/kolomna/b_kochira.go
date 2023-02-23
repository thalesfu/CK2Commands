package kolomna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科奇拉KochiraBarony struct {
	feud.BaseBarony
}

var BKochira科奇拉 feud.Barony = &科奇拉KochiraBarony{}

func init() {
	f := BKochira科奇拉.(*科奇拉KochiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kochira",
		TitleName: "科奇拉",
		TitleCode: "b_kochira",
	}
}
