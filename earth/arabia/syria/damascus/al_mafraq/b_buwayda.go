package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布未达BuwaydaBarony struct {
	feud.BaseBarony
}

var BBuwayda布未达 feud.Barony = &布未达BuwaydaBarony{}

func init() {
    f := BBuwayda布未达.(*布未达BuwaydaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buwayda",
		TitleName: "布未达",
		TitleCode: "b_buwayda",
	}
}
