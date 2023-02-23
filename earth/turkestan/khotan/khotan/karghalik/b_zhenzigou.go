package karghalik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 榛子沟ZhenzigouBarony struct {
	feud.BaseBarony
}

var BZhenzigou榛子沟 feud.Barony = &榛子沟ZhenzigouBarony{}

func init() {
	f := BZhenzigou榛子沟.(*榛子沟ZhenzigouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhenzigou",
		TitleName: "榛子沟",
		TitleCode: "b_zhenzigou",
	}
}
