package khokh_serkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔勒TalBarony struct {
	feud.BaseBarony
}

var BTal塔勒 feud.Barony = &塔勒TalBarony{}

func init() {
    f := BTal塔勒.(*塔勒TalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tal",
		TitleName: "塔勒",
		TitleCode: "b_tal",
	}
}
