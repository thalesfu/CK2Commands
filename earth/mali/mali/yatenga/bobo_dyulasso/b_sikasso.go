package bobo_dyulasso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西卡索SikassoBarony struct {
	feud.BaseBarony
}

var BSikasso西卡索 feud.Barony = &西卡索SikassoBarony{}

func init() {
    f := BSikasso西卡索.(*西卡索SikassoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sikasso",
		TitleName: "西卡索",
		TitleCode: "b_sikasso",
	}
}
