package zhongba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 仁多RingtorBarony struct {
	feud.BaseBarony
}

var BRingtor仁多 feud.Barony = &仁多RingtorBarony{}

func init() {
    f := BRingtor仁多.(*仁多RingtorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ringtor",
		TitleName: "仁多",
		TitleCode: "b_ringtor",
	}
}
