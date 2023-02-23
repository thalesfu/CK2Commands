package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬蒂尼PontignyBarony struct {
	feud.BaseBarony
}

var BPontigny蓬蒂尼 feud.Barony = &蓬蒂尼PontignyBarony{}

func init() {
	f := BPontigny蓬蒂尼.(*蓬蒂尼PontignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pontigny",
		TitleName: "蓬蒂尼",
		TitleCode: "b_pontigny",
	}
}
