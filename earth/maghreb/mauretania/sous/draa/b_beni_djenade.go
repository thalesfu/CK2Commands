package draa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尼杰纳德Beni_djenadeBarony struct {
	feud.BaseBarony
}

var BBeni_djenade贝尼杰纳德 feud.Barony = &贝尼杰纳德Beni_djenadeBarony{}

func init() {
    f := BBeni_djenade贝尼杰纳德.(*贝尼杰纳德Beni_djenadeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beni_djenade",
		TitleName: "贝尼杰纳德",
		TitleCode: "b_beni_djenade",
	}
}
