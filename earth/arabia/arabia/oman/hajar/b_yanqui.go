package hajar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 延古勒YanquiBarony struct {
	feud.BaseBarony
}

var BYanqui延古勒 feud.Barony = &延古勒YanquiBarony{}

func init() {
	f := BYanqui延古勒.(*延古勒YanquiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yanqui",
		TitleName: "延古勒",
		TitleCode: "b_yanqui",
	}
}
