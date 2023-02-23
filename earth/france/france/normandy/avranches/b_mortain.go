package avranches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔坦MortainBarony struct {
	feud.BaseBarony
}

var BMortain莫尔坦 feud.Barony = &莫尔坦MortainBarony{}

func init() {
	f := BMortain莫尔坦.(*莫尔坦MortainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mortain",
		TitleName: "莫尔坦",
		TitleCode: "b_mortain",
	}
}
