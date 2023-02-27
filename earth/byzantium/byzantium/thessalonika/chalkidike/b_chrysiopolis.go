package chalkidike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫律索波利斯ChrysiopolisBarony struct {
	feud.BaseBarony
}

var BChrysiopolis赫律索波利斯 feud.Barony = &赫律索波利斯ChrysiopolisBarony{}

func init() {
    f := BChrysiopolis赫律索波利斯.(*赫律索波利斯ChrysiopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chrysiopolis",
		TitleName: "赫律索波利斯",
		TitleCode: "b_chrysiopolis",
	}
}
