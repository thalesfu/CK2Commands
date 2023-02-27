package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布罗萨斯BrozasBarony struct {
	feud.BaseBarony
}

var BBrozas布罗萨斯 feud.Barony = &布罗萨斯BrozasBarony{}

func init() {
    f := BBrozas布罗萨斯.(*布罗萨斯BrozasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brozas",
		TitleName: "布罗萨斯",
		TitleCode: "b_brozas",
	}
}
