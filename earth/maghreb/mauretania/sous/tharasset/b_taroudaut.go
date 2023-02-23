package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔鲁丹特TaroudautBarony struct {
	feud.BaseBarony
}

var BTaroudaut塔鲁丹特 feud.Barony = &塔鲁丹特TaroudautBarony{}

func init() {
	f := BTaroudaut塔鲁丹特.(*塔鲁丹特TaroudautBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taroudaut",
		TitleName: "塔鲁丹特",
		TitleCode: "b_taroudaut",
	}
}
