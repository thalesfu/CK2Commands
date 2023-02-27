package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐姆良斯科耶TsimlyanskoyeBarony struct {
	feud.BaseBarony
}

var BTsimlyanskoye齐姆良斯科耶 feud.Barony = &齐姆良斯科耶TsimlyanskoyeBarony{}

func init() {
    f := BTsimlyanskoye齐姆良斯科耶.(*齐姆良斯科耶TsimlyanskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsimlyanskoye",
		TitleName: "齐姆良斯科耶",
		TitleCode: "b_tsimlyanskoye",
	}
}
