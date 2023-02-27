package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦隆AvallonBarony struct {
	feud.BaseBarony
}

var BAvallon阿瓦隆 feud.Barony = &阿瓦隆AvallonBarony{}

func init() {
    f := BAvallon阿瓦隆.(*阿瓦隆AvallonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avallon",
		TitleName: "阿瓦隆",
		TitleCode: "b_avallon",
	}
}
