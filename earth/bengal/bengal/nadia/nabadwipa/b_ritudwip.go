package nabadwipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利堵洲RitudwipBarony struct {
	feud.BaseBarony
}

var BRitudwip利堵洲 feud.Barony = &利堵洲RitudwipBarony{}

func init() {
	f := BRitudwip利堵洲.(*利堵洲RitudwipBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ritudwip",
		TitleName: "利堵洲",
		TitleCode: "b_ritudwip",
	}
}
