package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈利梅尔_尤KhalmeryuBarony struct {
	feud.BaseBarony
}

var BKhalmeryu哈利梅尔_尤 feud.Barony = &哈利梅尔_尤KhalmeryuBarony{}

func init() {
    f := BKhalmeryu哈利梅尔_尤.(*哈利梅尔_尤KhalmeryuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khalmeryu",
		TitleName: "哈利梅尔_尤",
		TitleCode: "b_khalmeryu",
	}
}
