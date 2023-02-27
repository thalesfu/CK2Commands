package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克套AqtawBarony struct {
	feud.BaseBarony
}

var BAqtaw阿克套 feud.Barony = &阿克套AqtawBarony{}

func init() {
    f := BAqtaw阿克套.(*阿克套AqtawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aqtaw",
		TitleName: "阿克套",
		TitleCode: "b_aqtaw",
	}
}
