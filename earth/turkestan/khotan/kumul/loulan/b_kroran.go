package loulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楼兰KroranBarony struct {
	feud.BaseBarony
}

var BKroran楼兰 feud.Barony = &楼兰KroranBarony{}

func init() {
    f := BKroran楼兰.(*楼兰KroranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kroran",
		TitleName: "楼兰",
		TitleCode: "b_kroran",
	}
}
