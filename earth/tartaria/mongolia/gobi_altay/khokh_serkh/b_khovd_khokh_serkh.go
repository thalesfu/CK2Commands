package khokh_serkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科布多Khovd_khokh_serkhBarony struct {
	feud.BaseBarony
}

var BKhovd_khokh_serkh科布多 feud.Barony = &科布多Khovd_khokh_serkhBarony{}

func init() {
    f := BKhovd_khokh_serkh科布多.(*科布多Khovd_khokh_serkhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khovd_khokh_serkh",
		TitleName: "科布多",
		TitleCode: "b_khovd_khokh_serkh",
	}
}
