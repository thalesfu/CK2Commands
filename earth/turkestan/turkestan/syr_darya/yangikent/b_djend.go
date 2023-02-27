package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毡的DjendBarony struct {
	feud.BaseBarony
}

var BDjend毡的 feud.Barony = &毡的DjendBarony{}

func init() {
    f := BDjend毡的.(*毡的DjendBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djend",
		TitleName: "毡的",
		TitleCode: "b_djend",
	}
}
