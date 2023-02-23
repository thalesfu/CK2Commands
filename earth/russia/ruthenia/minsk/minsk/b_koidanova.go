package minsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科伊达诺瓦KoidanovaBarony struct {
	feud.BaseBarony
}

var BKoidanova科伊达诺瓦 feud.Barony = &科伊达诺瓦KoidanovaBarony{}

func init() {
	f := BKoidanova科伊达诺瓦.(*科伊达诺瓦KoidanovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koidanova",
		TitleName: "科伊达诺瓦",
		TitleCode: "b_koidanova",
	}
}
