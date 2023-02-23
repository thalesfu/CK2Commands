package tsakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 聂尔NyerBarony struct {
	feud.BaseBarony
}

var BNyer聂尔 feud.Barony = &聂尔NyerBarony{}

func init() {
	f := BNyer聂尔.(*聂尔NyerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyer",
		TitleName: "聂尔",
		TitleCode: "b_nyer",
	}
}
