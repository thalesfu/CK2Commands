package tukums

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉BullaBarony struct {
	feud.BaseBarony
}

var BBulla布拉 feud.Barony = &布拉BullaBarony{}

func init() {
    f := BBulla布拉.(*布拉BullaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bulla",
		TitleName: "布拉",
		TitleCode: "b_bulla",
	}
}
