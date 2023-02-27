package chach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新城NavekatBarony struct {
	feud.BaseBarony
}

var BNavekat新城 feud.Barony = &新城NavekatBarony{}

func init() {
    f := BNavekat新城.(*新城NavekatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "navekat",
		TitleName: "新城",
		TitleCode: "b_navekat",
	}
}
