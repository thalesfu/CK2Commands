package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔门德KormendBarony struct {
	feud.BaseBarony
}

var BKormend克尔门德 feud.Barony = &克尔门德KormendBarony{}

func init() {
	f := BKormend克尔门德.(*克尔门德KormendBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kormend",
		TitleName: "克尔门德",
		TitleCode: "b_kormend",
	}
}
