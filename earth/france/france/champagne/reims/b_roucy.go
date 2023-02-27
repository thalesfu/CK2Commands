package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁西RoucyBarony struct {
	feud.BaseBarony
}

var BRoucy鲁西 feud.Barony = &鲁西RoucyBarony{}

func init() {
    f := BRoucy鲁西.(*鲁西RoucyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roucy",
		TitleName: "鲁西",
		TitleCode: "b_roucy",
	}
}
