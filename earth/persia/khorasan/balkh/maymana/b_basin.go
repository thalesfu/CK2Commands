package maymana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝欣BasinBarony struct {
	feud.BaseBarony
}

var BBasin贝欣 feud.Barony = &贝欣BasinBarony{}

func init() {
	f := BBasin贝欣.(*贝欣BasinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basin",
		TitleName: "贝欣",
		TitleCode: "b_basin",
	}
}
