package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多瑙沃特DonauworthBarony struct {
	feud.BaseBarony
}

var BDonauworth多瑙沃特 feud.Barony = &多瑙沃特DonauworthBarony{}

func init() {
    f := BDonauworth多瑙沃特.(*多瑙沃特DonauworthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donauworth",
		TitleName: "多瑙沃特",
		TitleCode: "b_donauworth",
	}
}
