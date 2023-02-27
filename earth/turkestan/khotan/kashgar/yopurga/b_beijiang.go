package yopurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 北降BeijiangBarony struct {
	feud.BaseBarony
}

var BBeijiang北降 feud.Barony = &北降BeijiangBarony{}

func init() {
    f := BBeijiang北降.(*北降BeijiangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beijiang",
		TitleName: "北降",
		TitleCode: "b_beijiang",
	}
}
