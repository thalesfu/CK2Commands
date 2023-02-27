package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔雷TaryBarony struct {
	feud.BaseBarony
}

var BTary塔雷 feud.Barony = &塔雷TaryBarony{}

func init() {
    f := BTary塔雷.(*塔雷TaryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tary",
		TitleName: "塔雷",
		TitleCode: "b_tary",
	}
}
