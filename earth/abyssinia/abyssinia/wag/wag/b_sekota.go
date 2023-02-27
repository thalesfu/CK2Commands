package wag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞科塔SekotaBarony struct {
	feud.BaseBarony
}

var BSekota塞科塔 feud.Barony = &塞科塔SekotaBarony{}

func init() {
    f := BSekota塞科塔.(*塞科塔SekotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sekota",
		TitleName: "塞科塔",
		TitleCode: "b_sekota",
	}
}
