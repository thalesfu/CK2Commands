package evreux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利雪LisieuxBarony struct {
	feud.BaseBarony
}

var BLisieux利雪 feud.Barony = &利雪LisieuxBarony{}

func init() {
    f := BLisieux利雪.(*利雪LisieuxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lisieux",
		TitleName: "利雪",
		TitleCode: "b_lisieux",
	}
}
