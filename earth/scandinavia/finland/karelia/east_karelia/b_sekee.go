package east_karelia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞凯SekeeBarony struct {
	feud.BaseBarony
}

var BSekee塞凯 feud.Barony = &塞凯SekeeBarony{}

func init() {
    f := BSekee塞凯.(*塞凯SekeeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sekee",
		TitleName: "塞凯",
		TitleCode: "b_sekee",
	}
}
