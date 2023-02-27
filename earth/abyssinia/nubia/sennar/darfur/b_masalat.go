package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马萨拉特MasalatBarony struct {
	feud.BaseBarony
}

var BMasalat马萨拉特 feud.Barony = &马萨拉特MasalatBarony{}

func init() {
    f := BMasalat马萨拉特.(*马萨拉特MasalatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masalat",
		TitleName: "马萨拉特",
		TitleCode: "b_masalat",
	}
}
