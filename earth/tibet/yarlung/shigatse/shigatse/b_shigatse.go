package shigatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日喀则ShigatseBarony struct {
	feud.BaseBarony
}

var BShigatse日喀则 feud.Barony = &日喀则ShigatseBarony{}

func init() {
	f := BShigatse日喀则.(*日喀则ShigatseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shigatse",
		TitleName: "日喀则",
		TitleCode: "b_shigatse",
	}
}
