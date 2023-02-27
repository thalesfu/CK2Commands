package severin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅哈迪亚MehadiaBarony struct {
	feud.BaseBarony
}

var BMehadia梅哈迪亚 feud.Barony = &梅哈迪亚MehadiaBarony{}

func init() {
    f := BMehadia梅哈迪亚.(*梅哈迪亚MehadiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mehadia",
		TitleName: "梅哈迪亚",
		TitleCode: "b_mehadia",
	}
}
