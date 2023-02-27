package kiranapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乾乔梨KankauliBarony struct {
	feud.BaseBarony
}

var BKankauli乾乔梨 feud.Barony = &乾乔梨KankauliBarony{}

func init() {
    f := BKankauli乾乔梨.(*乾乔梨KankauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kankauli",
		TitleName: "乾乔梨",
		TitleCode: "b_kankauli",
	}
}
