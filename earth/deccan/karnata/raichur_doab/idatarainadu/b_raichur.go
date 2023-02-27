package idatarainadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗耶注卢RaichurBarony struct {
	feud.BaseBarony
}

var BRaichur罗耶注卢 feud.Barony = &罗耶注卢RaichurBarony{}

func init() {
    f := BRaichur罗耶注卢.(*罗耶注卢RaichurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raichur",
		TitleName: "罗耶注卢",
		TitleCode: "b_raichur",
	}
}
