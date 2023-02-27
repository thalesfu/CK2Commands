package otrar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车尔纳克ChernakBarony struct {
	feud.BaseBarony
}

var BChernak车尔纳克 feud.Barony = &车尔纳克ChernakBarony{}

func init() {
    f := BChernak车尔纳克.(*车尔纳克ChernakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernak",
		TitleName: "车尔纳克",
		TitleCode: "b_chernak",
	}
}
