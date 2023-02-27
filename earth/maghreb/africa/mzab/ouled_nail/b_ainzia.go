package ouled_nail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因兹AinziaBarony struct {
	feud.BaseBarony
}

var BAinzia艾因兹 feud.Barony = &艾因兹AinziaBarony{}

func init() {
    f := BAinzia艾因兹.(*艾因兹AinziaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainzia",
		TitleName: "艾因兹",
		TitleCode: "b_ainzia",
	}
}
