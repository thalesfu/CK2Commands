package peresechen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁特RautuBarony struct {
	feud.BaseBarony
}

var BRautu鲁特 feud.Barony = &鲁特RautuBarony{}

func init() {
	f := BRautu鲁特.(*鲁特RautuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rautu",
		TitleName: "鲁特",
		TitleCode: "b_rautu",
	}
}
