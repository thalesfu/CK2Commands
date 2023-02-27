package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆陀BhaddarBarony struct {
	feud.BaseBarony
}

var BBhaddar婆陀 feud.Barony = &婆陀BhaddarBarony{}

func init() {
    f := BBhaddar婆陀.(*婆陀BhaddarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhaddar",
		TitleName: "婆陀",
		TitleCode: "b_bhaddar",
	}
}
