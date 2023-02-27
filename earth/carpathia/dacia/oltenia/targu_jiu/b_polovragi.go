package targu_jiu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波洛夫拉吉PolovragiBarony struct {
	feud.BaseBarony
}

var BPolovragi波洛夫拉吉 feud.Barony = &波洛夫拉吉PolovragiBarony{}

func init() {
    f := BPolovragi波洛夫拉吉.(*波洛夫拉吉PolovragiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polovragi",
		TitleName: "波洛夫拉吉",
		TitleCode: "b_polovragi",
	}
}
