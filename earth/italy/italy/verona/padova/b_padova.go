package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕多瓦PadovaBarony struct {
	feud.BaseBarony
}

var BPadova帕多瓦 feud.Barony = &帕多瓦PadovaBarony{}

func init() {
	f := BPadova帕多瓦.(*帕多瓦PadovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "padova",
		TitleName: "帕多瓦",
		TitleCode: "b_padova",
	}
}
