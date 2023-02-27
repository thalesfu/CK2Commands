package kolguyev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔基纳KonkinaBarony struct {
	feud.BaseBarony
}

var BKonkina孔基纳 feud.Barony = &孔基纳KonkinaBarony{}

func init() {
    f := BKonkina孔基纳.(*孔基纳KonkinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konkina",
		TitleName: "孔基纳",
		TitleCode: "b_konkina",
	}
}
