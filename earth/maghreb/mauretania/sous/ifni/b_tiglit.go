package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梯格利特TiglitBarony struct {
	feud.BaseBarony
}

var BTiglit梯格利特 feud.Barony = &梯格利特TiglitBarony{}

func init() {
	f := BTiglit梯格利特.(*梯格利特TiglitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiglit",
		TitleName: "梯格利特",
		TitleCode: "b_tiglit",
	}
}
