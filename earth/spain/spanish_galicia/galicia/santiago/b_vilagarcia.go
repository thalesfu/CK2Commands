package santiago

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比拉加尔西亚VilagarciaBarony struct {
	feud.BaseBarony
}

var BVilagarcia比拉加尔西亚 feud.Barony = &比拉加尔西亚VilagarciaBarony{}

func init() {
    f := BVilagarcia比拉加尔西亚.(*比拉加尔西亚VilagarciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vilagarcia",
		TitleName: "比拉加尔西亚",
		TitleCode: "b_vilagarcia",
	}
}
