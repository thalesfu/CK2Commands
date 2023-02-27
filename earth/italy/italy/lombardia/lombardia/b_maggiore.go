package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马焦雷MaggioreBarony struct {
	feud.BaseBarony
}

var BMaggiore马焦雷 feud.Barony = &马焦雷MaggioreBarony{}

func init() {
    f := BMaggiore马焦雷.(*马焦雷MaggioreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maggiore",
		TitleName: "马焦雷",
		TitleCode: "b_maggiore",
	}
}
