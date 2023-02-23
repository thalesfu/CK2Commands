package betpaqdala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅纳拉尔MynaralBarony struct {
	feud.BaseBarony
}

var BMynaral梅纳拉尔 feud.Barony = &梅纳拉尔MynaralBarony{}

func init() {
	f := BMynaral梅纳拉尔.(*梅纳拉尔MynaralBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mynaral",
		TitleName: "梅纳拉尔",
		TitleCode: "b_mynaral",
	}
}
