package mstislavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗波伊斯克PropoyskBarony struct {
	feud.BaseBarony
}

var BPropoysk普罗波伊斯克 feud.Barony = &普罗波伊斯克PropoyskBarony{}

func init() {
    f := BPropoysk普罗波伊斯克.(*普罗波伊斯克PropoyskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "propoysk",
		TitleName: "普罗波伊斯克",
		TitleCode: "b_propoysk",
	}
}
