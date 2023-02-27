package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉斯ArrasBarony struct {
	feud.BaseBarony
}

var BArras阿拉斯 feud.Barony = &阿拉斯ArrasBarony{}

func init() {
    f := BArras阿拉斯.(*阿拉斯ArrasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arras",
		TitleName: "阿拉斯",
		TitleCode: "b_arras",
	}
}
