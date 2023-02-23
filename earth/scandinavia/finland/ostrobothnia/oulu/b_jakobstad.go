package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮耶塔尔萨里JakobstadBarony struct {
	feud.BaseBarony
}

var BJakobstad皮耶塔尔萨里 feud.Barony = &皮耶塔尔萨里JakobstadBarony{}

func init() {
	f := BJakobstad皮耶塔尔萨里.(*皮耶塔尔萨里JakobstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jakobstad",
		TitleName: "皮耶塔尔萨里",
		TitleCode: "b_jakobstad",
	}
}
