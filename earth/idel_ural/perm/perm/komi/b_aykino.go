package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾基诺AykinoBarony struct {
	feud.BaseBarony
}

var BAykino艾基诺 feud.Barony = &艾基诺AykinoBarony{}

func init() {
    f := BAykino艾基诺.(*艾基诺AykinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aykino",
		TitleName: "艾基诺",
		TitleCode: "b_aykino",
	}
}
