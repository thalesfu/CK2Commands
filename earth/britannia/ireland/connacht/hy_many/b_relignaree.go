package hy_many

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷利格纳里RelignareeBarony struct {
	feud.BaseBarony
}

var BRelignaree雷利格纳里 feud.Barony = &雷利格纳里RelignareeBarony{}

func init() {
    f := BRelignaree雷利格纳里.(*雷利格纳里RelignareeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "relignaree",
		TitleName: "雷利格纳里",
		TitleCode: "b_relignaree",
	}
}
