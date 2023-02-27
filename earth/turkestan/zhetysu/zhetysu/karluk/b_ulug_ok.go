package karluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡禄屋Ulug_okBarony struct {
	feud.BaseBarony
}

var BUlug_ok胡禄屋 feud.Barony = &胡禄屋Ulug_okBarony{}

func init() {
    f := BUlug_ok胡禄屋.(*胡禄屋Ulug_okBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulug_ok",
		TitleName: "胡禄屋",
		TitleCode: "b_ulug_ok",
	}
}
