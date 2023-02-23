package chieti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿泰萨AtessaBarony struct {
	feud.BaseBarony
}

var BAtessa阿泰萨 feud.Barony = &阿泰萨AtessaBarony{}

func init() {
	f := BAtessa阿泰萨.(*阿泰萨AtessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atessa",
		TitleName: "阿泰萨",
		TitleCode: "b_atessa",
	}
}
