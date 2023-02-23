package ujjayini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆思DewasBarony struct {
	feud.BaseBarony
}

var BDewas提婆思 feud.Barony = &提婆思DewasBarony{}

func init() {
	f := BDewas提婆思.(*提婆思DewasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dewas",
		TitleName: "提婆思",
		TitleCode: "b_dewas",
	}
}
