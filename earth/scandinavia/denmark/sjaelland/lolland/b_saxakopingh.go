package lolland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨克斯克宾SaxakopinghBarony struct {
	feud.BaseBarony
}

var BSaxakopingh萨克斯克宾 feud.Barony = &萨克斯克宾SaxakopinghBarony{}

func init() {
    f := BSaxakopingh萨克斯克宾.(*萨克斯克宾SaxakopinghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saxakopingh",
		TitleName: "萨克斯克宾",
		TitleCode: "b_saxakopingh",
	}
}
