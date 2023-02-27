package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃吉斯海姆EgisheimBarony struct {
	feud.BaseBarony
}

var BEgisheim埃吉斯海姆 feud.Barony = &埃吉斯海姆EgisheimBarony{}

func init() {
    f := BEgisheim埃吉斯海姆.(*埃吉斯海姆EgisheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egisheim",
		TitleName: "埃吉斯海姆",
		TitleCode: "b_egisheim",
	}
}
