package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿琴泽拉AcenzeraBarony struct {
	feud.BaseBarony
}

var BAcenzera阿琴泽拉 feud.Barony = &阿琴泽拉AcenzeraBarony{}

func init() {
	f := BAcenzera阿琴泽拉.(*阿琴泽拉AcenzeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acenzera",
		TitleName: "阿琴泽拉",
		TitleCode: "b_acenzera",
	}
}
