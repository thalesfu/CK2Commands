package khovsgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥诺特OnotBarony struct {
	feud.BaseBarony
}

var BOnot奥诺特 feud.Barony = &奥诺特OnotBarony{}

func init() {
	f := BOnot奥诺特.(*奥诺特OnotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "onot",
		TitleName: "奥诺特",
		TitleCode: "b_onot",
	}
}
