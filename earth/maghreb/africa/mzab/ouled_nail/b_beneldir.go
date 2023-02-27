package ouled_nail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本迪尔BeneldirBarony struct {
	feud.BaseBarony
}

var BBeneldir本迪尔 feud.Barony = &本迪尔BeneldirBarony{}

func init() {
    f := BBeneldir本迪尔.(*本迪尔BeneldirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beneldir",
		TitleName: "本迪尔",
		TitleCode: "b_beneldir",
	}
}
