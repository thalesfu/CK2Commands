package shahrazur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆卡勒AmkalehBarony struct {
	feud.BaseBarony
}

var BAmkaleh阿姆卡勒 feud.Barony = &阿姆卡勒AmkalehBarony{}

func init() {
	f := BAmkaleh阿姆卡勒.(*阿姆卡勒AmkalehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amkaleh",
		TitleName: "阿姆卡勒",
		TitleCode: "b_amkaleh",
	}
}
