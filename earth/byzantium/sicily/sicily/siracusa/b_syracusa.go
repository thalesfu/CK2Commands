package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡拉库萨SyracusaBarony struct {
	feud.BaseBarony
}

var BSyracusa锡拉库萨 feud.Barony = &锡拉库萨SyracusaBarony{}

func init() {
	f := BSyracusa锡拉库萨.(*锡拉库萨SyracusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syracusa",
		TitleName: "锡拉库萨",
		TitleCode: "b_syracusa",
	}
}
