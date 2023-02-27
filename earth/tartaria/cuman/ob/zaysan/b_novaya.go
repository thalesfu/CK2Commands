package zaysan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺瓦亚NovayaBarony struct {
	feud.BaseBarony
}

var BNovaya诺瓦亚 feud.Barony = &诺瓦亚NovayaBarony{}

func init() {
    f := BNovaya诺瓦亚.(*诺瓦亚NovayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novaya",
		TitleName: "诺瓦亚",
		TitleCode: "b_novaya",
	}
}
