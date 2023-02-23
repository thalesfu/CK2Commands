package korosten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢吉尼LouhynyBarony struct {
	feud.BaseBarony
}

var BLouhyny卢吉尼 feud.Barony = &卢吉尼LouhynyBarony{}

func init() {
	f := BLouhyny卢吉尼.(*卢吉尼LouhynyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "louhyny",
		TitleName: "卢吉尼",
		TitleCode: "b_louhyny",
	}
}
