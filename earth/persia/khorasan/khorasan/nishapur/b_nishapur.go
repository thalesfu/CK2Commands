package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内沙布尔NishapurBarony struct {
	feud.BaseBarony
}

var BNishapur内沙布尔 feud.Barony = &内沙布尔NishapurBarony{}

func init() {
    f := BNishapur内沙布尔.(*内沙布尔NishapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nishapur",
		TitleName: "内沙布尔",
		TitleCode: "b_nishapur",
	}
}
