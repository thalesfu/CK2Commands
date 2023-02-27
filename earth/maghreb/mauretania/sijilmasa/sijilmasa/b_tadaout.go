package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔达乌特TadaoutBarony struct {
	feud.BaseBarony
}

var BTadaout塔达乌特 feud.Barony = &塔达乌特TadaoutBarony{}

func init() {
    f := BTadaout塔达乌特.(*塔达乌特TadaoutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tadaout",
		TitleName: "塔达乌特",
		TitleCode: "b_tadaout",
	}
}
