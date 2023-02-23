package nyland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈米纳HaminaBarony struct {
	feud.BaseBarony
}

var BHamina哈米纳 feud.Barony = &哈米纳HaminaBarony{}

func init() {
	f := BHamina哈米纳.(*哈米纳HaminaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamina",
		TitleName: "哈米纳",
		TitleCode: "b_hamina",
	}
}
