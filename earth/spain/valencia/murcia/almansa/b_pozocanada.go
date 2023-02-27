package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波索卡尼亚达PozocanadaBarony struct {
	feud.BaseBarony
}

var BPozocanada波索卡尼亚达 feud.Barony = &波索卡尼亚达PozocanadaBarony{}

func init() {
    f := BPozocanada波索卡尼亚达.(*波索卡尼亚达PozocanadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pozocanada",
		TitleName: "波索卡尼亚达",
		TitleCode: "b_pozocanada",
	}
}
