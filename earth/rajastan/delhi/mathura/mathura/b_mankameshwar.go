package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩那迦迷湿伐罗MankameshwarBarony struct {
	feud.BaseBarony
}

var BMankameshwar摩那迦迷湿伐罗 feud.Barony = &摩那迦迷湿伐罗MankameshwarBarony{}

func init() {
	f := BMankameshwar摩那迦迷湿伐罗.(*摩那迦迷湿伐罗MankameshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mankameshwar",
		TitleName: "摩那迦迷湿伐罗",
		TitleCode: "b_mankameshwar",
	}
}
