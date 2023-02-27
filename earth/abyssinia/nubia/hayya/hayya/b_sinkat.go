package hayya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛卡特SinkatBarony struct {
	feud.BaseBarony
}

var BSinkat辛卡特 feud.Barony = &辛卡特SinkatBarony{}

func init() {
    f := BSinkat辛卡特.(*辛卡特SinkatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sinkat",
		TitleName: "辛卡特",
		TitleCode: "b_sinkat",
	}
}
