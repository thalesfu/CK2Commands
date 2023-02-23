package iarmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉利TraleeBarony struct {
	feud.BaseBarony
}

var BTralee特拉利 feud.Barony = &特拉利TraleeBarony{}

func init() {
	f := BTralee特拉利.(*特拉利TraleeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tralee",
		TitleName: "特拉利",
		TitleCode: "b_tralee",
	}
}
