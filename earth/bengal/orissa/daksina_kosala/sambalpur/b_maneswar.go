package sambalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩泥湿伐罗ManeswarBarony struct {
	feud.BaseBarony
}

var BManeswar摩泥湿伐罗 feud.Barony = &摩泥湿伐罗ManeswarBarony{}

func init() {
    f := BManeswar摩泥湿伐罗.(*摩泥湿伐罗ManeswarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maneswar",
		TitleName: "摩泥湿伐罗",
		TitleCode: "b_maneswar",
	}
}
