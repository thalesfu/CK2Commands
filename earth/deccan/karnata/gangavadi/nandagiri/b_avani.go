package nandagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦尼AvaniBarony struct {
	feud.BaseBarony
}

var BAvani阿瓦尼 feud.Barony = &阿瓦尼AvaniBarony{}

func init() {
	f := BAvani阿瓦尼.(*阿瓦尼AvaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avani",
		TitleName: "阿瓦尼",
		TitleCode: "b_avani",
	}
}
