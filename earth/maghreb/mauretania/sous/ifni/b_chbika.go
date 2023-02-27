package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍比凯ChbikaBarony struct {
	feud.BaseBarony
}

var BChbika舍比凯 feud.Barony = &舍比凯ChbikaBarony{}

func init() {
    f := BChbika舍比凯.(*舍比凯ChbikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chbika",
		TitleName: "舍比凯",
		TitleCode: "b_chbika",
	}
}
