package sarakhs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 撒剌哈夕SarakhsBarony struct {
	feud.BaseBarony
}

var BSarakhs撒剌哈夕 feud.Barony = &撒剌哈夕SarakhsBarony{}

func init() {
	f := BSarakhs撒剌哈夕.(*撒剌哈夕SarakhsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarakhs",
		TitleName: "撒剌哈夕",
		TitleCode: "b_sarakhs",
	}
}
