package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希埃隆HieronBarony struct {
	feud.BaseBarony
}

var BHieron希埃隆 feud.Barony = &希埃隆HieronBarony{}

func init() {
	f := BHieron希埃隆.(*希埃隆HieronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hieron",
		TitleName: "希埃隆",
		TitleCode: "b_hieron",
	}
}
