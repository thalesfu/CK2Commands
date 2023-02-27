package joensuu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约恩苏JoensuuBarony struct {
	feud.BaseBarony
}

var BJoensuu约恩苏 feud.Barony = &约恩苏JoensuuBarony{}

func init() {
    f := BJoensuu约恩苏.(*约恩苏JoensuuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "joensuu",
		TitleName: "约恩苏",
		TitleCode: "b_joensuu",
	}
}
