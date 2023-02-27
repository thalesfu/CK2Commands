package kalaus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔马维尔ArmavirBarony struct {
	feud.BaseBarony
}

var BArmavir阿尔马维尔 feud.Barony = &阿尔马维尔ArmavirBarony{}

func init() {
    f := BArmavir阿尔马维尔.(*阿尔马维尔ArmavirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "armavir",
		TitleName: "阿尔马维尔",
		TitleCode: "b_armavir",
	}
}
