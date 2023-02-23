package chilia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂盖奇TigheciBarony struct {
	feud.BaseBarony
}

var BTigheci蒂盖奇 feud.Barony = &蒂盖奇TigheciBarony{}

func init() {
	f := BTigheci蒂盖奇.(*蒂盖奇TigheciBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigheci",
		TitleName: "蒂盖奇",
		TitleCode: "b_tigheci",
	}
}
