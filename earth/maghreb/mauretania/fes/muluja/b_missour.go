package muluja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米苏尔MissourBarony struct {
	feud.BaseBarony
}

var BMissour米苏尔 feud.Barony = &米苏尔MissourBarony{}

func init() {
	f := BMissour米苏尔.(*米苏尔MissourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "missour",
		TitleName: "米苏尔",
		TitleCode: "b_missour",
	}
}
