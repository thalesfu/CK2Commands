package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科普里扬KoprijanBarony struct {
	feud.BaseBarony
}

var BKoprijan科普里扬 feud.Barony = &科普里扬KoprijanBarony{}

func init() {
	f := BKoprijan科普里扬.(*科普里扬KoprijanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koprijan",
		TitleName: "科普里扬",
		TitleCode: "b_koprijan",
	}
}
