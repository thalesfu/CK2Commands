package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉比苏斯ArabissusBarony struct {
	feud.BaseBarony
}

var BArabissus阿拉比苏斯 feud.Barony = &阿拉比苏斯ArabissusBarony{}

func init() {
    f := BArabissus阿拉比苏斯.(*阿拉比苏斯ArabissusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arabissus",
		TitleName: "阿拉比苏斯",
		TitleCode: "b_arabissus",
	}
}
