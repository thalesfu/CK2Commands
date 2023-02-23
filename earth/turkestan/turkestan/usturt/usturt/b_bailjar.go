package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜伊加尔BailjarBarony struct {
	feud.BaseBarony
}

var BBailjar拜伊加尔 feud.Barony = &拜伊加尔BailjarBarony{}

func init() {
	f := BBailjar拜伊加尔.(*拜伊加尔BailjarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bailjar",
		TitleName: "拜伊加尔",
		TitleCode: "b_bailjar",
	}
}
