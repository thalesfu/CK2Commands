package tanggula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛曲MaquBarony struct {
	feud.BaseBarony
}

var BMaqu玛曲 feud.Barony = &玛曲MaquBarony{}

func init() {
    f := BMaqu玛曲.(*玛曲MaquBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maqu",
		TitleName: "玛曲",
		TitleCode: "b_maqu",
	}
}
