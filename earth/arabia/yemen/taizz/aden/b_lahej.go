package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉哈杰LahejBarony struct {
	feud.BaseBarony
}

var BLahej拉哈杰 feud.Barony = &拉哈杰LahejBarony{}

func init() {
	f := BLahej拉哈杰.(*拉哈杰LahejBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lahej",
		TitleName: "拉哈杰",
		TitleCode: "b_lahej",
	}
}
