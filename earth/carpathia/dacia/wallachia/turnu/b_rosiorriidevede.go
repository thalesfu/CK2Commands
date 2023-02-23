package turnu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗肖里德韦代RosiorriidevedeBarony struct {
	feud.BaseBarony
}

var BRosiorriidevede罗肖里德韦代 feud.Barony = &罗肖里德韦代RosiorriidevedeBarony{}

func init() {
	f := BRosiorriidevede罗肖里德韦代.(*罗肖里德韦代RosiorriidevedeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rosiorriidevede",
		TitleName: "罗肖里德韦代",
		TitleCode: "b_rosiorriidevede",
	}
}
