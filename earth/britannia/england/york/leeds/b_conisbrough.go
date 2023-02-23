package leeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尼斯伯勒ConisbroughBarony struct {
	feud.BaseBarony
}

var BConisbrough科尼斯伯勒 feud.Barony = &科尼斯伯勒ConisbroughBarony{}

func init() {
	f := BConisbrough科尼斯伯勒.(*科尼斯伯勒ConisbroughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "conisbrough",
		TitleName: "科尼斯伯勒",
		TitleCode: "b_conisbrough",
	}
}
