package gabiyaha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布托ButoBarony struct {
	feud.BaseBarony
}

var BButo布托 feud.Barony = &布托ButoBarony{}

func init() {
	f := BButo布托.(*布托ButoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buto",
		TitleName: "布托",
		TitleCode: "b_buto",
	}
}
