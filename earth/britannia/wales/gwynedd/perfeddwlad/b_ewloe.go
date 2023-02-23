package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤洛EwloeBarony struct {
	feud.BaseBarony
}

var BEwloe尤洛 feud.Barony = &尤洛EwloeBarony{}

func init() {
	f := BEwloe尤洛.(*尤洛EwloeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ewloe",
		TitleName: "尤洛",
		TitleCode: "b_ewloe",
	}
}
