package kyzyl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿扎尔加AzargaBarony struct {
	feud.BaseBarony
}

var BAzarga阿扎尔加 feud.Barony = &阿扎尔加AzargaBarony{}

func init() {
    f := BAzarga阿扎尔加.(*阿扎尔加AzargaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azarga",
		TitleName: "阿扎尔加",
		TitleCode: "b_azarga",
	}
}
