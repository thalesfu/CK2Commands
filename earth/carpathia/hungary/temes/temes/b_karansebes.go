package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考兰舍贝什KaransebesBarony struct {
	feud.BaseBarony
}

var BKaransebes考兰舍贝什 feud.Barony = &考兰舍贝什KaransebesBarony{}

func init() {
	f := BKaransebes考兰舍贝什.(*考兰舍贝什KaransebesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karansebes",
		TitleName: "考兰舍贝什",
		TitleCode: "b_karansebes",
	}
}
