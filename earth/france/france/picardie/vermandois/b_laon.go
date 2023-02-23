package vermandois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉昂LaonBarony struct {
	feud.BaseBarony
}

var BLaon拉昂 feud.Barony = &拉昂LaonBarony{}

func init() {
	f := BLaon拉昂.(*拉昂LaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laon",
		TitleName: "拉昂",
		TitleCode: "b_laon",
	}
}
