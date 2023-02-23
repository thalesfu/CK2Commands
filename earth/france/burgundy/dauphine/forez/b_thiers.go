package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梯也尔ThiersBarony struct {
	feud.BaseBarony
}

var BThiers梯也尔 feud.Barony = &梯也尔ThiersBarony{}

func init() {
	f := BThiers梯也尔.(*梯也尔ThiersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thiers",
		TitleName: "梯也尔",
		TitleCode: "b_thiers",
	}
}
