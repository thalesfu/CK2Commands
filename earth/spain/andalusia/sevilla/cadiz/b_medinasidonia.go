package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡多尼亚城MedinasidoniaBarony struct {
	feud.BaseBarony
}

var BMedinasidonia锡多尼亚城 feud.Barony = &锡多尼亚城MedinasidoniaBarony{}

func init() {
	f := BMedinasidonia锡多尼亚城.(*锡多尼亚城MedinasidoniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medinasidonia",
		TitleName: "锡多尼亚城",
		TitleCode: "b_medinasidonia",
	}
}
