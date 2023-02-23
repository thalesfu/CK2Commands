package dasapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀卢DaluBarony struct {
	feud.BaseBarony
}

var BDalu陀卢 feud.Barony = &陀卢DaluBarony{}

func init() {
	f := BDalu陀卢.(*陀卢DaluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dalu",
		TitleName: "陀卢",
		TitleCode: "b_dalu",
	}
}
