package oppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗莱斯贝格FlesbergBarony struct {
	feud.BaseBarony
}

var BFlesberg弗莱斯贝格 feud.Barony = &弗莱斯贝格FlesbergBarony{}

func init() {
	f := BFlesberg弗莱斯贝格.(*弗莱斯贝格FlesbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "flesberg",
		TitleName: "弗莱斯贝格",
		TitleCode: "b_flesberg",
	}
}
