package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布扎奇亚拉AbuzqayqaBarony struct {
	feud.BaseBarony
}

var BAbuzqayqa阿布扎奇亚拉 feud.Barony = &阿布扎奇亚拉AbuzqayqaBarony{}

func init() {
	f := BAbuzqayqa阿布扎奇亚拉.(*阿布扎奇亚拉AbuzqayqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abuzqayqa",
		TitleName: "阿布扎奇亚拉",
		TitleCode: "b_abuzqayqa",
	}
}
