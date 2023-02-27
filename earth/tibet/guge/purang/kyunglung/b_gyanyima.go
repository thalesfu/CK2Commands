package kyunglung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姜叶玛GyanyimaBarony struct {
	feud.BaseBarony
}

var BGyanyima姜叶玛 feud.Barony = &姜叶玛GyanyimaBarony{}

func init() {
    f := BGyanyima姜叶玛.(*姜叶玛GyanyimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyanyima",
		TitleName: "姜叶玛",
		TitleCode: "b_gyanyima",
	}
}
