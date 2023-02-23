package tara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克伊泽斯KeyzesBarony struct {
	feud.BaseBarony
}

var BKeyzes克伊泽斯 feud.Barony = &克伊泽斯KeyzesBarony{}

func init() {
	f := BKeyzes克伊泽斯.(*克伊泽斯KeyzesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keyzes",
		TitleName: "克伊泽斯",
		TitleCode: "b_keyzes",
	}
}
