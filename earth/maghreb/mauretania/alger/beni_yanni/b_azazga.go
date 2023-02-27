package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿扎兹加AzazgaBarony struct {
	feud.BaseBarony
}

var BAzazga阿扎兹加 feud.Barony = &阿扎兹加AzazgaBarony{}

func init() {
    f := BAzazga阿扎兹加.(*阿扎兹加AzazgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azazga",
		TitleName: "阿扎兹加",
		TitleCode: "b_azazga",
	}
}
