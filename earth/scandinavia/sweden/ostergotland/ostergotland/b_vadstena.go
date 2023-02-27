package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦斯泰纳VadstenaBarony struct {
	feud.BaseBarony
}

var BVadstena瓦斯泰纳 feud.Barony = &瓦斯泰纳VadstenaBarony{}

func init() {
    f := BVadstena瓦斯泰纳.(*瓦斯泰纳VadstenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vadstena",
		TitleName: "瓦斯泰纳",
		TitleCode: "b_vadstena",
	}
}
