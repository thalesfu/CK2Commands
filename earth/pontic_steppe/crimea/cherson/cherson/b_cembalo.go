package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琴巴洛CembaloBarony struct {
	feud.BaseBarony
}

var BCembalo琴巴洛 feud.Barony = &琴巴洛CembaloBarony{}

func init() {
    f := BCembalo琴巴洛.(*琴巴洛CembaloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cembalo",
		TitleName: "琴巴洛",
		TitleCode: "b_cembalo",
	}
}
