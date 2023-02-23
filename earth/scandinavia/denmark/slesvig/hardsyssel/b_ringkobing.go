package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 灵克宾RingkobingBarony struct {
	feud.BaseBarony
}

var BRingkobing灵克宾 feud.Barony = &灵克宾RingkobingBarony{}

func init() {
	f := BRingkobing灵克宾.(*灵克宾RingkobingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ringkobing",
		TitleName: "灵克宾",
		TitleCode: "b_ringkobing",
	}
}
