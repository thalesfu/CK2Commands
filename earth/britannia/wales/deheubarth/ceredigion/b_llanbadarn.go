package ceredigion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰巴达恩LlanbadarnBarony struct {
	feud.BaseBarony
}

var BLlanbadarn兰巴达恩 feud.Barony = &兰巴达恩LlanbadarnBarony{}

func init() {
	f := BLlanbadarn兰巴达恩.(*兰巴达恩LlanbadarnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanbadarn",
		TitleName: "兰巴达恩",
		TitleCode: "b_llanbadarn",
	}
}
