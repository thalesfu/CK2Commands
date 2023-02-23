package kirghiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛涅列SinneleBarony struct {
	feud.BaseBarony
}

var BSinnele辛涅列 feud.Barony = &辛涅列SinneleBarony{}

func init() {
	f := BSinnele辛涅列.(*辛涅列SinneleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sinnele",
		TitleName: "辛涅列",
		TitleCode: "b_sinnele",
	}
}
