package om

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博罗金卡BorodinkaBarony struct {
	feud.BaseBarony
}

var BBorodinka博罗金卡 feud.Barony = &博罗金卡BorodinkaBarony{}

func init() {
    f := BBorodinka博罗金卡.(*博罗金卡BorodinkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borodinka",
		TitleName: "博罗金卡",
		TitleCode: "b_borodinka",
	}
}
