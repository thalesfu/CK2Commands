package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因塔卜AintabBarony struct {
	feud.BaseBarony
}

var BAintab艾因塔卜 feud.Barony = &艾因塔卜AintabBarony{}

func init() {
    f := BAintab艾因塔卜.(*艾因塔卜AintabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aintab",
		TitleName: "艾因塔卜",
		TitleCode: "b_aintab",
	}
}
