package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马斯拉法尔MathrafalBarony struct {
	feud.BaseBarony
}

var BMathrafal马斯拉法尔 feud.Barony = &马斯拉法尔MathrafalBarony{}

func init() {
	f := BMathrafal马斯拉法尔.(*马斯拉法尔MathrafalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mathrafal",
		TitleName: "马斯拉法尔",
		TitleCode: "b_mathrafal",
	}
}
