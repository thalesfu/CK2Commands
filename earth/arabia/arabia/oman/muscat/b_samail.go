package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞马伊勒SamailBarony struct {
	feud.BaseBarony
}

var BSamail塞马伊勒 feud.Barony = &塞马伊勒SamailBarony{}

func init() {
    f := BSamail塞马伊勒.(*塞马伊勒SamailBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samail",
		TitleName: "塞马伊勒",
		TitleCode: "b_samail",
	}
}
