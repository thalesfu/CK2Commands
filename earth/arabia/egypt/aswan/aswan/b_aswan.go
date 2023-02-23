package aswan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯旺AswanBarony struct {
	feud.BaseBarony
}

var BAswan阿斯旺 feud.Barony = &阿斯旺AswanBarony{}

func init() {
	f := BAswan阿斯旺.(*阿斯旺AswanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aswan",
		TitleName: "阿斯旺",
		TitleCode: "b_aswan",
	}
}
