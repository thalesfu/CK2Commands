package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮尼利亚PinillaBarony struct {
	feud.BaseBarony
}

var BPinilla皮尼利亚 feud.Barony = &皮尼利亚PinillaBarony{}

func init() {
    f := BPinilla皮尼利亚.(*皮尼利亚PinillaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pinilla",
		TitleName: "皮尼利亚",
		TitleCode: "b_pinilla",
	}
}
