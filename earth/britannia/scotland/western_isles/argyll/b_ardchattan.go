package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾彻坦ArdchattanBarony struct {
	feud.BaseBarony
}

var BArdchattan艾彻坦 feud.Barony = &艾彻坦ArdchattanBarony{}

func init() {
    f := BArdchattan艾彻坦.(*艾彻坦ArdchattanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardchattan",
		TitleName: "艾彻坦",
		TitleCode: "b_ardchattan",
	}
}
