package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔玛拉TamraBarony struct {
	feud.BaseBarony
}

var BTamra塔玛拉 feud.Barony = &塔玛拉TamraBarony{}

func init() {
    f := BTamra塔玛拉.(*塔玛拉TamraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamra",
		TitleName: "塔玛拉",
		TitleCode: "b_tamra",
	}
}
