package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希格拉HigraBarony struct {
	feud.BaseBarony
}

var BHigra希格拉 feud.Barony = &希格拉HigraBarony{}

func init() {
    f := BHigra希格拉.(*希格拉HigraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "higra",
		TitleName: "希格拉",
		TitleCode: "b_higra",
	}
}
