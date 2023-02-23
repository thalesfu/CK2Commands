package tara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉TaraBarony struct {
	feud.BaseBarony
}

var BTara塔拉 feud.Barony = &塔拉TaraBarony{}

func init() {
	f := BTara塔拉.(*塔拉TaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tara",
		TitleName: "塔拉",
		TitleCode: "b_tara",
	}
}
