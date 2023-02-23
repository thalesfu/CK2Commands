package bezichy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克柳切瓦亚KlyuchevayaBarony struct {
	feud.BaseBarony
}

var BKlyuchevaya克柳切瓦亚 feud.Barony = &克柳切瓦亚KlyuchevayaBarony{}

func init() {
	f := BKlyuchevaya克柳切瓦亚.(*克柳切瓦亚KlyuchevayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klyuchevaya",
		TitleName: "克柳切瓦亚",
		TitleCode: "b_klyuchevaya",
	}
}
