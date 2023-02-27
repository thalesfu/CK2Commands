package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚喀巴AqabahBarony struct {
	feud.BaseBarony
}

var BAqabah亚喀巴 feud.Barony = &亚喀巴AqabahBarony{}

func init() {
    f := BAqabah亚喀巴.(*亚喀巴AqabahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aqabah",
		TitleName: "亚喀巴",
		TitleCode: "b_aqabah",
	}
}
