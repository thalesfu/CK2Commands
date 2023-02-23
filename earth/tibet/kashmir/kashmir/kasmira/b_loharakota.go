package kasmira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢诃罗拘吒LoharakotaBarony struct {
	feud.BaseBarony
}

var BLoharakota卢诃罗拘吒 feud.Barony = &卢诃罗拘吒LoharakotaBarony{}

func init() {
	f := BLoharakota卢诃罗拘吒.(*卢诃罗拘吒LoharakotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loharakota",
		TitleName: "卢诃罗拘吒",
		TitleCode: "b_loharakota",
	}
}
