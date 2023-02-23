package tao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马姆罗瓦尼MamrovaniBarony struct {
	feud.BaseBarony
}

var BMamrovani马姆罗瓦尼 feud.Barony = &马姆罗瓦尼MamrovaniBarony{}

func init() {
	f := BMamrovani马姆罗瓦尼.(*马姆罗瓦尼MamrovaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamrovani",
		TitleName: "马姆罗瓦尼",
		TitleCode: "b_mamrovani",
	}
}
