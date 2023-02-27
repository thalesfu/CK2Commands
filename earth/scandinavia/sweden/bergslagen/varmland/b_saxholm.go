package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨克斯霍尔姆SaxholmBarony struct {
	feud.BaseBarony
}

var BSaxholm萨克斯霍尔姆 feud.Barony = &萨克斯霍尔姆SaxholmBarony{}

func init() {
    f := BSaxholm萨克斯霍尔姆.(*萨克斯霍尔姆SaxholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saxholm",
		TitleName: "萨克斯霍尔姆",
		TitleCode: "b_saxholm",
	}
}
