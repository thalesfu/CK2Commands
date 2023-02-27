package berbera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔奥BurcoBarony struct {
	feud.BaseBarony
}

var BBurco布尔奥 feud.Barony = &布尔奥BurcoBarony{}

func init() {
    f := BBurco布尔奥.(*布尔奥BurcoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burco",
		TitleName: "布尔奥",
		TitleCode: "b_burco",
	}
}
