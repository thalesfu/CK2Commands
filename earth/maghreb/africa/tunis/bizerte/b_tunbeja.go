package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴杰TunbejaBarony struct {
	feud.BaseBarony
}

var BTunbeja巴杰 feud.Barony = &巴杰TunbejaBarony{}

func init() {
	f := BTunbeja巴杰.(*巴杰TunbejaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tunbeja",
		TitleName: "巴杰",
		TitleCode: "b_tunbeja",
	}
}
