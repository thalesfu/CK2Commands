package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈纳咖XonqaBarony struct {
	feud.BaseBarony
}

var BXonqa哈纳咖 feud.Barony = &哈纳咖XonqaBarony{}

func init() {
    f := BXonqa哈纳咖.(*哈纳咖XonqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xonqa",
		TitleName: "哈纳咖",
		TitleCode: "b_xonqa",
	}
}
