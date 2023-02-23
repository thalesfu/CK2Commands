package mansura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内兰乔特NerunkotBarony struct {
	feud.BaseBarony
}

var BNerunkot内兰乔特 feud.Barony = &内兰乔特NerunkotBarony{}

func init() {
	f := BNerunkot内兰乔特.(*内兰乔特NerunkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nerunkot",
		TitleName: "内兰乔特",
		TitleCode: "b_nerunkot",
	}
}
