package caen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔VireBarony struct {
	feud.BaseBarony
}

var BVire维尔 feud.Barony = &维尔VireBarony{}

func init() {
	f := BVire维尔.(*维尔VireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vire",
		TitleName: "维尔",
		TitleCode: "b_vire",
	}
}
