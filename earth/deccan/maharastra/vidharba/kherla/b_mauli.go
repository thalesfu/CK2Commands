package kherla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩梨MauliBarony struct {
	feud.BaseBarony
}

var BMauli摩梨 feud.Barony = &摩梨MauliBarony{}

func init() {
	f := BMauli摩梨.(*摩梨MauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mauli",
		TitleName: "摩梨",
		TitleCode: "b_mauli",
	}
}
