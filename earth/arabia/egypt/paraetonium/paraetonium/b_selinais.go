package paraetonium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞利涅斯SelinaisBarony struct {
	feud.BaseBarony
}

var BSelinais塞利涅斯 feud.Barony = &塞利涅斯SelinaisBarony{}

func init() {
	f := BSelinais塞利涅斯.(*塞利涅斯SelinaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selinais",
		TitleName: "塞利涅斯",
		TitleCode: "b_selinais",
	}
}
