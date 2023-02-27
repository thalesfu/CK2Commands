package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊维IviosBarony struct {
	feud.BaseBarony
}

var BIvios伊维 feud.Barony = &伊维IviosBarony{}

func init() {
    f := BIvios伊维.(*伊维IviosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivios",
		TitleName: "伊维",
		TitleCode: "b_ivios",
	}
}
