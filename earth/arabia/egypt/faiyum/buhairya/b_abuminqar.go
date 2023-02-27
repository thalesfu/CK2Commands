package buhairya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布民加尔AbuminqarBarony struct {
	feud.BaseBarony
}

var BAbuminqar阿布民加尔 feud.Barony = &阿布民加尔AbuminqarBarony{}

func init() {
    f := BAbuminqar阿布民加尔.(*阿布民加尔AbuminqarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abuminqar",
		TitleName: "阿布民加尔",
		TitleCode: "b_abuminqar",
	}
}
