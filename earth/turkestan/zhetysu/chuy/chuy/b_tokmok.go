package chuy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托克马克TokmokBarony struct {
	feud.BaseBarony
}

var BTokmok托克马克 feud.Barony = &托克马克TokmokBarony{}

func init() {
	f := BTokmok托克马克.(*托克马克TokmokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tokmok",
		TitleName: "托克马克",
		TitleCode: "b_tokmok",
	}
}
