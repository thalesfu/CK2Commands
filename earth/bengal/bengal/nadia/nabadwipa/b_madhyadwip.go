package nabadwipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩地耶洲MadhyadwipBarony struct {
	feud.BaseBarony
}

var BMadhyadwip摩地耶洲 feud.Barony = &摩地耶洲MadhyadwipBarony{}

func init() {
    f := BMadhyadwip摩地耶洲.(*摩地耶洲MadhyadwipBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madhyadwip",
		TitleName: "摩地耶洲",
		TitleCode: "b_madhyadwip",
	}
}
