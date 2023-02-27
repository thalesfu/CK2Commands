package arigh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里格ArighBarony struct {
	feud.BaseBarony
}

var BArigh阿里格 feud.Barony = &阿里格ArighBarony{}

func init() {
    f := BArigh阿里格.(*阿里格ArighBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arigh",
		TitleName: "阿里格",
		TitleCode: "b_arigh",
	}
}
