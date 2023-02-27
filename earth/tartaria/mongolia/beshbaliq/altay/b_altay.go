package altay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔泰AltayBarony struct {
	feud.BaseBarony
}

var BAltay阿尔泰 feud.Barony = &阿尔泰AltayBarony{}

func init() {
    f := BAltay阿尔泰.(*阿尔泰AltayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altay",
		TitleName: "阿尔泰",
		TitleCode: "b_altay",
	}
}
