package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯库涅斯SkulnasBarony struct {
	feud.BaseBarony
}

var BSkulnas斯库涅斯 feud.Barony = &斯库涅斯SkulnasBarony{}

func init() {
    f := BSkulnas斯库涅斯.(*斯库涅斯SkulnasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skulnas",
		TitleName: "斯库涅斯",
		TitleCode: "b_skulnas",
	}
}
