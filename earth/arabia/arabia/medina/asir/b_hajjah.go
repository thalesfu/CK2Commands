package asir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈杰HajjahBarony struct {
	feud.BaseBarony
}

var BHajjah哈杰 feud.Barony = &哈杰HajjahBarony{}

func init() {
	f := BHajjah哈杰.(*哈杰HajjahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hajjah",
		TitleName: "哈杰",
		TitleCode: "b_hajjah",
	}
}
