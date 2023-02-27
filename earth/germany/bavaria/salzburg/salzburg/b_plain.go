package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普莱恩PlainBarony struct {
	feud.BaseBarony
}

var BPlain普莱恩 feud.Barony = &普莱恩PlainBarony{}

func init() {
    f := BPlain普莱恩.(*普莱恩PlainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plain",
		TitleName: "普莱恩",
		TitleCode: "b_plain",
	}
}
