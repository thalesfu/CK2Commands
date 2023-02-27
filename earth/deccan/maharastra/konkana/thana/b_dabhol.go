package thana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达布尔DabholBarony struct {
	feud.BaseBarony
}

var BDabhol达布尔 feud.Barony = &达布尔DabholBarony{}

func init() {
    f := BDabhol达布尔.(*达布尔DabholBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dabhol",
		TitleName: "达布尔",
		TitleCode: "b_dabhol",
	}
}
