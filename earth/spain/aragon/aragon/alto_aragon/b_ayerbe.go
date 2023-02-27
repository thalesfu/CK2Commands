package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿耶韦AyerbeBarony struct {
	feud.BaseBarony
}

var BAyerbe阿耶韦 feud.Barony = &阿耶韦AyerbeBarony{}

func init() {
    f := BAyerbe阿耶韦.(*阿耶韦AyerbeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayerbe",
		TitleName: "阿耶韦",
		TitleCode: "b_ayerbe",
	}
}
