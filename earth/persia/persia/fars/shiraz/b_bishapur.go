package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比沙普尔BishapurBarony struct {
	feud.BaseBarony
}

var BBishapur比沙普尔 feud.Barony = &比沙普尔BishapurBarony{}

func init() {
	f := BBishapur比沙普尔.(*比沙普尔BishapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bishapur",
		TitleName: "比沙普尔",
		TitleCode: "b_bishapur",
	}
}
