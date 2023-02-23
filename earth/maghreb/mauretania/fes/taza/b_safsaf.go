package taza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞夫萨夫SafsafBarony struct {
	feud.BaseBarony
}

var BSafsaf塞夫萨夫 feud.Barony = &塞夫萨夫SafsafBarony{}

func init() {
	f := BSafsaf塞夫萨夫.(*塞夫萨夫SafsafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "safsaf",
		TitleName: "塞夫萨夫",
		TitleCode: "b_safsaf",
	}
}
