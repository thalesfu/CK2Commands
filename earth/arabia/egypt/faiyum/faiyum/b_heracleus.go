package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉克琉斯HeracleusBarony struct {
	feud.BaseBarony
}

var BHeracleus赫拉克琉斯 feud.Barony = &赫拉克琉斯HeracleusBarony{}

func init() {
	f := BHeracleus赫拉克琉斯.(*赫拉克琉斯HeracleusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heracleus",
		TitleName: "赫拉克琉斯",
		TitleCode: "b_heracleus",
	}
}
