package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬斯韦克拉赫Pons_veckelaheBarony struct {
	feud.BaseBarony
}

var BPons_veckelahe蓬斯韦克拉赫 feud.Barony = &蓬斯韦克拉赫Pons_veckelaheBarony{}

func init() {
    f := BPons_veckelahe蓬斯韦克拉赫.(*蓬斯韦克拉赫Pons_veckelaheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pons_veckelahe",
		TitleName: "蓬斯韦克拉赫",
		TitleCode: "b_pons_veckelahe",
	}
}
