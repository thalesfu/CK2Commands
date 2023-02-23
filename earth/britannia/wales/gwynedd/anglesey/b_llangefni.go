package anglesey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰盖夫尼LlangefniBarony struct {
	feud.BaseBarony
}

var BLlangefni兰盖夫尼 feud.Barony = &兰盖夫尼LlangefniBarony{}

func init() {
	f := BLlangefni兰盖夫尼.(*兰盖夫尼LlangefniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llangefni",
		TitleName: "兰盖夫尼",
		TitleCode: "b_llangefni",
	}
}
