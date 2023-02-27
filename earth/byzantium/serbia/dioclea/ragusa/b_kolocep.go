package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科罗切普KolocepBarony struct {
	feud.BaseBarony
}

var BKolocep科罗切普 feud.Barony = &科罗切普KolocepBarony{}

func init() {
    f := BKolocep科罗切普.(*科罗切普KolocepBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolocep",
		TitleName: "科罗切普",
		TitleCode: "b_kolocep",
	}
}
