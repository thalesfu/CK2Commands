package medantaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波伽PakalBarony struct {
	feud.BaseBarony
}

var BPakal波伽 feud.Barony = &波伽PakalBarony{}

func init() {
	f := BPakal波伽.(*波伽PakalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pakal",
		TitleName: "波伽",
		TitleCode: "b_pakal",
	}
}
