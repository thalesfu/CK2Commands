package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯凯恩SkjernBarony struct {
	feud.BaseBarony
}

var BSkjern斯凯恩 feud.Barony = &斯凯恩SkjernBarony{}

func init() {
    f := BSkjern斯凯恩.(*斯凯恩SkjernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skjern",
		TitleName: "斯凯恩",
		TitleCode: "b_skjern",
	}
}
