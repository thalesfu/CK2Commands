package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 腓特烈港FriedrichshafenBarony struct {
	feud.BaseBarony
}

var BFriedrichshafen腓特烈港 feud.Barony = &腓特烈港FriedrichshafenBarony{}

func init() {
	f := BFriedrichshafen腓特烈港.(*腓特烈港FriedrichshafenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "friedrichshafen",
		TitleName: "腓特烈港",
		TitleCode: "b_friedrichshafen",
	}
}
