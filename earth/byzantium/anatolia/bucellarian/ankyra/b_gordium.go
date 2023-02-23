package ankyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔迪乌姆GordiumBarony struct {
	feud.BaseBarony
}

var BGordium戈尔迪乌姆 feud.Barony = &戈尔迪乌姆GordiumBarony{}

func init() {
	f := BGordium戈尔迪乌姆.(*戈尔迪乌姆GordiumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gordium",
		TitleName: "戈尔迪乌姆",
		TitleCode: "b_gordium",
	}
}
