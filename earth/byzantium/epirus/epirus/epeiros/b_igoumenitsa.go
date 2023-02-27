package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊古迈尼察IgoumenitsaBarony struct {
	feud.BaseBarony
}

var BIgoumenitsa伊古迈尼察 feud.Barony = &伊古迈尼察IgoumenitsaBarony{}

func init() {
    f := BIgoumenitsa伊古迈尼察.(*伊古迈尼察IgoumenitsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "igoumenitsa",
		TitleName: "伊古迈尼察",
		TitleCode: "b_igoumenitsa",
	}
}
