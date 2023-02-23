package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大陶波尔恰尼NagytapolcsanyBarony struct {
	feud.BaseBarony
}

var BNagytapolcsany大陶波尔恰尼 feud.Barony = &大陶波尔恰尼NagytapolcsanyBarony{}

func init() {
	f := BNagytapolcsany大陶波尔恰尼.(*大陶波尔恰尼NagytapolcsanyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagytapolcsany",
		TitleName: "大陶波尔恰尼",
		TitleCode: "b_nagytapolcsany",
	}
}
