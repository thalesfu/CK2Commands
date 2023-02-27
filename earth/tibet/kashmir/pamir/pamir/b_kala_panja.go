package pamir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀喇喷赤Kala_panjaBarony struct {
	feud.BaseBarony
}

var BKala_panja喀喇喷赤 feud.Barony = &喀喇喷赤Kala_panjaBarony{}

func init() {
    f := BKala_panja喀喇喷赤.(*喀喇喷赤Kala_panjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kala_panja",
		TitleName: "喀喇喷赤",
		TitleCode: "b_kala_panja",
	}
}
