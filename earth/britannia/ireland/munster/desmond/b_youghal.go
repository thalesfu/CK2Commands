package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约尔YoughalBarony struct {
	feud.BaseBarony
}

var BYoughal约尔 feud.Barony = &约尔YoughalBarony{}

func init() {
	f := BYoughal约尔.(*约尔YoughalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "youghal",
		TitleName: "约尔",
		TitleCode: "b_youghal",
	}
}
