package dwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马兹纳贝德MatsnaberdBarony struct {
	feud.BaseBarony
}

var BMatsnaberd马兹纳贝德 feud.Barony = &马兹纳贝德MatsnaberdBarony{}

func init() {
    f := BMatsnaberd马兹纳贝德.(*马兹纳贝德MatsnaberdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matsnaberd",
		TitleName: "马兹纳贝德",
		TitleCode: "b_matsnaberd",
	}
}
