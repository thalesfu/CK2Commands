package ajadabiya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧盖莱Al_aghailaBarony struct {
	feud.BaseBarony
}

var BAl_aghaila欧盖莱 feud.Barony = &欧盖莱Al_aghailaBarony{}

func init() {
    f := BAl_aghaila欧盖莱.(*欧盖莱Al_aghailaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_aghaila",
		TitleName: "欧盖莱",
		TitleCode: "b_al_aghaila",
	}
}
