package perugia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿西西AssisiBarony struct {
	feud.BaseBarony
}

var BAssisi阿西西 feud.Barony = &阿西西AssisiBarony{}

func init() {
	f := BAssisi阿西西.(*阿西西AssisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assisi",
		TitleName: "阿西西",
		TitleCode: "b_assisi",
	}
}
