package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔塔斯TartasBarony struct {
	feud.BaseBarony
}

var BTartas塔尔塔斯 feud.Barony = &塔尔塔斯TartasBarony{}

func init() {
	f := BTartas塔尔塔斯.(*塔尔塔斯TartasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tartas",
		TitleName: "塔尔塔斯",
		TitleCode: "b_tartas",
	}
}
