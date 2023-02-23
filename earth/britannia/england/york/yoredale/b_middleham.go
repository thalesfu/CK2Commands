package yoredale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米德尔赫姆MiddlehamBarony struct {
	feud.BaseBarony
}

var BMiddleham米德尔赫姆 feud.Barony = &米德尔赫姆MiddlehamBarony{}

func init() {
	f := BMiddleham米德尔赫姆.(*米德尔赫姆MiddlehamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "middleham",
		TitleName: "米德尔赫姆",
		TitleCode: "b_middleham",
	}
}
