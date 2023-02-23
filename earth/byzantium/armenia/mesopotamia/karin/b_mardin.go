package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔丁MardinBarony struct {
	feud.BaseBarony
}

var BMardin马尔丁 feud.Barony = &马尔丁MardinBarony{}

func init() {
	f := BMardin马尔丁.(*马尔丁MardinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mardin",
		TitleName: "马尔丁",
		TitleCode: "b_mardin",
	}
}
