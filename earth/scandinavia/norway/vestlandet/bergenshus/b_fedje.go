package bergenshus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费迪厄FedjeBarony struct {
	feud.BaseBarony
}

var BFedje费迪厄 feud.Barony = &费迪厄FedjeBarony{}

func init() {
	f := BFedje费迪厄.(*费迪厄FedjeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fedje",
		TitleName: "费迪厄",
		TitleCode: "b_fedje",
	}
}
