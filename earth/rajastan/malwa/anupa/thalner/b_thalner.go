package thalner

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 闼罗内ThalnerBarony struct {
	feud.BaseBarony
}

var BThalner闼罗内 feud.Barony = &闼罗内ThalnerBarony{}

func init() {
    f := BThalner闼罗内.(*闼罗内ThalnerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thalner",
		TitleName: "闼罗内",
		TitleCode: "b_thalner",
	}
}
