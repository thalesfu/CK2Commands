package mandapika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩醯湿伐罗MaheshwarBarony struct {
	feud.BaseBarony
}

var BMaheshwar摩醯湿伐罗 feud.Barony = &摩醯湿伐罗MaheshwarBarony{}

func init() {
	f := BMaheshwar摩醯湿伐罗.(*摩醯湿伐罗MaheshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maheshwar",
		TitleName: "摩醯湿伐罗",
		TitleCode: "b_maheshwar",
	}
}
