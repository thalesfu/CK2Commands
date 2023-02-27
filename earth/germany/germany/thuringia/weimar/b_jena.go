package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶拿JenaBarony struct {
	feud.BaseBarony
}

var BJena耶拿 feud.Barony = &耶拿JenaBarony{}

func init() {
    f := BJena耶拿.(*耶拿JenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jena",
		TitleName: "耶拿",
		TitleCode: "b_jena",
	}
}
