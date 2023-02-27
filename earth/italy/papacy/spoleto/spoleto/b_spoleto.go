package spoleto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯波莱托SpoletoBarony struct {
	feud.BaseBarony
}

var BSpoleto斯波莱托 feud.Barony = &斯波莱托SpoletoBarony{}

func init() {
    f := BSpoleto斯波莱托.(*斯波莱托SpoletoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spoleto",
		TitleName: "斯波莱托",
		TitleCode: "b_spoleto",
	}
}
