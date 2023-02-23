package ramagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 槃阇波梨BanjapalliBarony struct {
	feud.BaseBarony
}

var BBanjapalli槃阇波梨 feud.Barony = &槃阇波梨BanjapalliBarony{}

func init() {
	f := BBanjapalli槃阇波梨.(*槃阇波梨BanjapalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banjapalli",
		TitleName: "槃阇波梨",
		TitleCode: "b_banjapalli",
	}
}
