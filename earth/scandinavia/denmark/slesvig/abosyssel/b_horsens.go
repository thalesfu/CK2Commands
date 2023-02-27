package abosyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔森斯HorsensBarony struct {
	feud.BaseBarony
}

var BHorsens霍尔森斯 feud.Barony = &霍尔森斯HorsensBarony{}

func init() {
    f := BHorsens霍尔森斯.(*霍尔森斯HorsensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horsens",
		TitleName: "霍尔森斯",
		TitleCode: "b_horsens",
	}
}
