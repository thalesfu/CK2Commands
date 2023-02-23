package sibi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马斯坦迦MastanjBarony struct {
	feud.BaseBarony
}

var BMastanj马斯坦迦 feud.Barony = &马斯坦迦MastanjBarony{}

func init() {
	f := BMastanj马斯坦迦.(*马斯坦迦MastanjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mastanj",
		TitleName: "马斯坦迦",
		TitleCode: "b_mastanj",
	}
}
