package dendi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉BuraBarony struct {
	feud.BaseBarony
}

var BBura布拉 feud.Barony = &布拉BuraBarony{}

func init() {
    f := BBura布拉.(*布拉BuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bura",
		TitleName: "布拉",
		TitleCode: "b_bura",
	}
}
