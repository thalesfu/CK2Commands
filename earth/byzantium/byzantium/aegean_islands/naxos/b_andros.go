package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德罗斯AndrosBarony struct {
	feud.BaseBarony
}

var BAndros安德罗斯 feud.Barony = &安德罗斯AndrosBarony{}

func init() {
    f := BAndros安德罗斯.(*安德罗斯AndrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andros",
		TitleName: "安德罗斯",
		TitleCode: "b_andros",
	}
}
