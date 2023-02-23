package kalinganagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯陵伽那揭罗KalinganagaraBarony struct {
	feud.BaseBarony
}

var BKalinganagara羯陵伽那揭罗 feud.Barony = &羯陵伽那揭罗KalinganagaraBarony{}

func init() {
	f := BKalinganagara羯陵伽那揭罗.(*羯陵伽那揭罗KalinganagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalinganagara",
		TitleName: "羯陵伽那揭罗",
		TitleCode: "b_kalinganagara",
	}
}
