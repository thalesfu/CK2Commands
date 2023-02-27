package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱斯马黑戈LesmahagowBarony struct {
	feud.BaseBarony
}

var BLesmahagow莱斯马黑戈 feud.Barony = &莱斯马黑戈LesmahagowBarony{}

func init() {
    f := BLesmahagow莱斯马黑戈.(*莱斯马黑戈LesmahagowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lesmahagow",
		TitleName: "莱斯马黑戈",
		TitleCode: "b_lesmahagow",
	}
}
