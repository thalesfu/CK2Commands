package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔凯斯坦TakestanBarony struct {
	feud.BaseBarony
}

var BTakestan塔凯斯坦 feud.Barony = &塔凯斯坦TakestanBarony{}

func init() {
    f := BTakestan塔凯斯坦.(*塔凯斯坦TakestanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "takestan",
		TitleName: "塔凯斯坦",
		TitleCode: "b_takestan",
	}
}
