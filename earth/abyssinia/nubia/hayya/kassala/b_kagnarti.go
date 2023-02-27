package kassala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡格纳尔蒂KagnartiBarony struct {
	feud.BaseBarony
}

var BKagnarti卡格纳尔蒂 feud.Barony = &卡格纳尔蒂KagnartiBarony{}

func init() {
    f := BKagnarti卡格纳尔蒂.(*卡格纳尔蒂KagnartiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kagnarti",
		TitleName: "卡格纳尔蒂",
		TitleCode: "b_kagnarti",
	}
}
