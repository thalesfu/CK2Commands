package khokh_serkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陶勒包TolboBarony struct {
	feud.BaseBarony
}

var BTolbo陶勒包 feud.Barony = &陶勒包TolboBarony{}

func init() {
    f := BTolbo陶勒包.(*陶勒包TolboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tolbo",
		TitleName: "陶勒包",
		TitleCode: "b_tolbo",
	}
}
