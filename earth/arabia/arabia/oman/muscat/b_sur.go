package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏尔SurBarony struct {
	feud.BaseBarony
}

var BSur苏尔 feud.Barony = &苏尔SurBarony{}

func init() {
	f := BSur苏尔.(*苏尔SurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sur",
		TitleName: "苏尔",
		TitleCode: "b_sur",
	}
}
