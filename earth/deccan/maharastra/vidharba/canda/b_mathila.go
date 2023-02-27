package canda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩底陀MathilaBarony struct {
	feud.BaseBarony
}

var BMathila摩底陀 feud.Barony = &摩底陀MathilaBarony{}

func init() {
    f := BMathila摩底陀.(*摩底陀MathilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mathila",
		TitleName: "摩底陀",
		TitleCode: "b_mathila",
	}
}
