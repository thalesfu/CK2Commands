package badghis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布兹甘BuzganBarony struct {
	feud.BaseBarony
}

var BBuzgan布兹甘 feud.Barony = &布兹甘BuzganBarony{}

func init() {
    f := BBuzgan布兹甘.(*布兹甘BuzganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buzgan",
		TitleName: "布兹甘",
		TitleCode: "b_buzgan",
	}
}
