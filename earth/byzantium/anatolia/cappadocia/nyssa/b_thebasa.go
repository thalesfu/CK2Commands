package nyssa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞巴萨ThebasaBarony struct {
	feud.BaseBarony
}

var BThebasa塞巴萨 feud.Barony = &塞巴萨ThebasaBarony{}

func init() {
    f := BThebasa塞巴萨.(*塞巴萨ThebasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thebasa",
		TitleName: "塞巴萨",
		TitleCode: "b_thebasa",
	}
}
