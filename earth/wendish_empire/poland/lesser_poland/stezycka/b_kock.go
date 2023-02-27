package stezycka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科茨克KockBarony struct {
	feud.BaseBarony
}

var BKock科茨克 feud.Barony = &科茨克KockBarony{}

func init() {
    f := BKock科茨克.(*科茨克KockBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kock",
		TitleName: "科茨克",
		TitleCode: "b_kock",
	}
}
