package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨莫拉ZamoraBarony struct {
	feud.BaseBarony
}

var BZamora萨莫拉 feud.Barony = &萨莫拉ZamoraBarony{}

func init() {
    f := BZamora萨莫拉.(*萨莫拉ZamoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zamora",
		TitleName: "萨莫拉",
		TitleCode: "b_zamora",
	}
}
