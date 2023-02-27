package vel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武克特尔Vuktyl_velBarony struct {
	feud.BaseBarony
}

var BVuktyl_vel武克特尔 feud.Barony = &武克特尔Vuktyl_velBarony{}

func init() {
    f := BVuktyl_vel武克特尔.(*武克特尔Vuktyl_velBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vuktyl_vel",
		TitleName: "武克特尔",
		TitleCode: "b_vuktyl_vel",
	}
}
