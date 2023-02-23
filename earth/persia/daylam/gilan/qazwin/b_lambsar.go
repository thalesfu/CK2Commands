package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉莫萨尔LambsarBarony struct {
	feud.BaseBarony
}

var BLambsar拉莫萨尔 feud.Barony = &拉莫萨尔LambsarBarony{}

func init() {
	f := BLambsar拉莫萨尔.(*拉莫萨尔LambsarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lambsar",
		TitleName: "拉莫萨尔",
		TitleCode: "b_lambsar",
	}
}
