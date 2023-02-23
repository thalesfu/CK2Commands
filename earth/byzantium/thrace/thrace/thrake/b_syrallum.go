package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希拉冷SyrallumBarony struct {
	feud.BaseBarony
}

var BSyrallum希拉冷 feud.Barony = &希拉冷SyrallumBarony{}

func init() {
	f := BSyrallum希拉冷.(*希拉冷SyrallumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syrallum",
		TitleName: "希拉冷",
		TitleCode: "b_syrallum",
	}
}
