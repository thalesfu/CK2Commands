package mahoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽牟陀KamodhBarony struct {
	feud.BaseBarony
}

var BKamodh伽牟陀 feud.Barony = &伽牟陀KamodhBarony{}

func init() {
    f := BKamodh伽牟陀.(*伽牟陀KamodhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamodh",
		TitleName: "伽牟陀",
		TitleCode: "b_kamodh",
	}
}
