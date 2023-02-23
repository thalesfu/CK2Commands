package bharauli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆达利BhadriBarony struct {
	feud.BaseBarony
}

var BBhadri婆达利 feud.Barony = &婆达利BhadriBarony{}

func init() {
	f := BBhadri婆达利.(*婆达利BhadriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadri",
		TitleName: "婆达利",
		TitleCode: "b_bhadri",
	}
}
