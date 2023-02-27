package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃皮拉EpilaBarony struct {
	feud.BaseBarony
}

var BEpila埃皮拉 feud.Barony = &埃皮拉EpilaBarony{}

func init() {
    f := BEpila埃皮拉.(*埃皮拉EpilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "epila",
		TitleName: "埃皮拉",
		TitleCode: "b_epila",
	}
}
