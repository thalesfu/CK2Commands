package racakonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩罗MaraBarony struct {
	feud.BaseBarony
}

var BMara摩罗 feud.Barony = &摩罗MaraBarony{}

func init() {
	f := BMara摩罗.(*摩罗MaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mara",
		TitleName: "摩罗",
		TitleCode: "b_mara",
	}
}
