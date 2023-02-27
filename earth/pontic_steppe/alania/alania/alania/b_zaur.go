package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎乌尔ZaurBarony struct {
	feud.BaseBarony
}

var BZaur扎乌尔 feud.Barony = &扎乌尔ZaurBarony{}

func init() {
    f := BZaur扎乌尔.(*扎乌尔ZaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaur",
		TitleName: "扎乌尔",
		TitleCode: "b_zaur",
	}
}
