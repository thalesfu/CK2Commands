package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希兰KhinanBarony struct {
	feud.BaseBarony
}

var BKhinan希兰 feud.Barony = &希兰KhinanBarony{}

func init() {
    f := BKhinan希兰.(*希兰KhinanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khinan",
		TitleName: "希兰",
		TitleCode: "b_khinan",
	}
}
