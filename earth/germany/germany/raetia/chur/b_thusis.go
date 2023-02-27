package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图西斯ThusisBarony struct {
	feud.BaseBarony
}

var BThusis图西斯 feud.Barony = &图西斯ThusisBarony{}

func init() {
    f := BThusis图西斯.(*图西斯ThusisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thusis",
		TitleName: "图西斯",
		TitleCode: "b_thusis",
	}
}
