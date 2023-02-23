package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米兰MilanoBarony struct {
	feud.BaseBarony
}

var BMilano米兰 feud.Barony = &米兰MilanoBarony{}

func init() {
	f := BMilano米兰.(*米兰MilanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "milano",
		TitleName: "米兰",
		TitleCode: "b_milano",
	}
}
