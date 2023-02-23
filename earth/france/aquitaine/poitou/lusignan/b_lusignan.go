package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕西尼昂LusignanBarony struct {
	feud.BaseBarony
}

var BLusignan吕西尼昂 feud.Barony = &吕西尼昂LusignanBarony{}

func init() {
	f := BLusignan吕西尼昂.(*吕西尼昂LusignanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lusignan",
		TitleName: "吕西尼昂",
		TitleCode: "b_lusignan",
	}
}
