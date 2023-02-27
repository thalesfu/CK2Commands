package opava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥帕瓦OpavaBarony struct {
	feud.BaseBarony
}

var BOpava奥帕瓦 feud.Barony = &奥帕瓦OpavaBarony{}

func init() {
    f := BOpava奥帕瓦.(*奥帕瓦OpavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "opava",
		TitleName: "奥帕瓦",
		TitleCode: "b_opava",
	}
}
