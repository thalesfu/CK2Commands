package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利默里克LimerickBarony struct {
	feud.BaseBarony
}

var BLimerick利默里克 feud.Barony = &利默里克LimerickBarony{}

func init() {
	f := BLimerick利默里克.(*利默里克LimerickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "limerick",
		TitleName: "利默里克",
		TitleCode: "b_limerick",
	}
}
