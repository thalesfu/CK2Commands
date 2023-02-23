package ajayameru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗瓦哈LawahBarony struct {
	feud.BaseBarony
}

var BLawah罗瓦哈 feud.Barony = &罗瓦哈LawahBarony{}

func init() {
	f := BLawah罗瓦哈.(*罗瓦哈LawahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lawah",
		TitleName: "罗瓦哈",
		TitleCode: "b_lawah",
	}
}
