package purushapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙_吉_谢里Shah_ji_dheriBarony struct {
	feud.BaseBarony
}

var BShah_ji_dheri沙_吉_谢里 feud.Barony = &沙_吉_谢里Shah_ji_dheriBarony{}

func init() {
    f := BShah_ji_dheri沙_吉_谢里.(*沙_吉_谢里Shah_ji_dheriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shah_ji_dheri",
		TitleName: "沙_吉_谢里",
		TitleCode: "b_shah_ji_dheri",
	}
}
