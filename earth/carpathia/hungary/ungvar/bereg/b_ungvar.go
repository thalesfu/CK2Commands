package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温格瓦尔UngvarBarony struct {
	feud.BaseBarony
}

var BUngvar温格瓦尔 feud.Barony = &温格瓦尔UngvarBarony{}

func init() {
	f := BUngvar温格瓦尔.(*温格瓦尔UngvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ungvar",
		TitleName: "温格瓦尔",
		TitleCode: "b_ungvar",
	}
}
