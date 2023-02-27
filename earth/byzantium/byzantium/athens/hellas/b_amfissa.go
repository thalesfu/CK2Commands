package hellas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆菲萨AmfissaBarony struct {
	feud.BaseBarony
}

var BAmfissa阿姆菲萨 feud.Barony = &阿姆菲萨AmfissaBarony{}

func init() {
    f := BAmfissa阿姆菲萨.(*阿姆菲萨AmfissaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amfissa",
		TitleName: "阿姆菲萨",
		TitleCode: "b_amfissa",
	}
}
