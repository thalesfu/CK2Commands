package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德维尼察DvinitsaBarony struct {
	feud.BaseBarony
}

var BDvinitsa德维尼察 feud.Barony = &德维尼察DvinitsaBarony{}

func init() {
	f := BDvinitsa德维尼察.(*德维尼察DvinitsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dvinitsa",
		TitleName: "德维尼察",
		TitleCode: "b_dvinitsa",
	}
}
