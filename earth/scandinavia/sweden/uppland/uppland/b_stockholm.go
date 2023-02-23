package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯德哥尔摩StockholmBarony struct {
	feud.BaseBarony
}

var BStockholm斯德哥尔摩 feud.Barony = &斯德哥尔摩StockholmBarony{}

func init() {
	f := BStockholm斯德哥尔摩.(*斯德哥尔摩StockholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stockholm",
		TitleName: "斯德哥尔摩",
		TitleCode: "b_stockholm",
	}
}
