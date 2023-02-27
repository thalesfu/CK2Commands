package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达罗摩汗那尤DairodmorhannanyoBarony struct {
	feud.BaseBarony
}

var BDairodmorhannanyo达罗摩汗那尤 feud.Barony = &达罗摩汗那尤DairodmorhannanyoBarony{}

func init() {
    f := BDairodmorhannanyo达罗摩汗那尤.(*达罗摩汗那尤DairodmorhannanyoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dairodmorhannanyo",
		TitleName: "达罗摩汗那尤",
		TitleCode: "b_dairodmorhannanyo",
	}
}
