package artux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿图什AtushBarony struct {
	feud.BaseBarony
}

var BAtush阿图什 feud.Barony = &阿图什AtushBarony{}

func init() {
	f := BAtush阿图什.(*阿图什AtushBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atush",
		TitleName: "阿图什",
		TitleCode: "b_atush",
	}
}
