package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏尔塞SurseeBarony struct {
	feud.BaseBarony
}

var BSursee苏尔塞 feud.Barony = &苏尔塞SurseeBarony{}

func init() {
    f := BSursee苏尔塞.(*苏尔塞SurseeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sursee",
		TitleName: "苏尔塞",
		TitleCode: "b_sursee",
	}
}
