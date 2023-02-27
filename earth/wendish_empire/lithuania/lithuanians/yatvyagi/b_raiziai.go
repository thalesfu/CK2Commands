package yatvyagi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖日艾RaiziaiBarony struct {
	feud.BaseBarony
}

var BRaiziai赖日艾 feud.Barony = &赖日艾RaiziaiBarony{}

func init() {
    f := BRaiziai赖日艾.(*赖日艾RaiziaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raiziai",
		TitleName: "赖日艾",
		TitleCode: "b_raiziai",
	}
}
