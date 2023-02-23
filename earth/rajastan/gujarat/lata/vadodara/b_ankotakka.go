package vadodara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安拘吒迦AnkotakkaBarony struct {
	feud.BaseBarony
}

var BAnkotakka安拘吒迦 feud.Barony = &安拘吒迦AnkotakkaBarony{}

func init() {
	f := BAnkotakka安拘吒迦.(*安拘吒迦AnkotakkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ankotakka",
		TitleName: "安拘吒迦",
		TitleCode: "b_ankotakka",
	}
}
