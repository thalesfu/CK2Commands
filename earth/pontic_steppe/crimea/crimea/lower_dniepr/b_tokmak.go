package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托克马克TokmakBarony struct {
	feud.BaseBarony
}

var BTokmak托克马克 feud.Barony = &托克马克TokmakBarony{}

func init() {
    f := BTokmak托克马克.(*托克马克TokmakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tokmak",
		TitleName: "托克马克",
		TitleCode: "b_tokmak",
	}
}
