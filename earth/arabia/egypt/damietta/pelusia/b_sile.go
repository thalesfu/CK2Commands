package pelusia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞勒SileBarony struct {
	feud.BaseBarony
}

var BSile塞勒 feud.Barony = &塞勒SileBarony{}

func init() {
    f := BSile塞勒.(*塞勒SileBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sile",
		TitleName: "塞勒",
		TitleCode: "b_sile",
	}
}
