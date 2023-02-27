package tis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂斯TisBarony struct {
	feud.BaseBarony
}

var BTis蒂斯 feud.Barony = &蒂斯TisBarony{}

func init() {
    f := BTis蒂斯.(*蒂斯TisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tis",
		TitleName: "蒂斯",
		TitleCode: "b_tis",
	}
}
