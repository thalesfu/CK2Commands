package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波斯拉BozraBarony struct {
	feud.BaseBarony
}

var BBozra波斯拉 feud.Barony = &波斯拉BozraBarony{}

func init() {
    f := BBozra波斯拉.(*波斯拉BozraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bozra",
		TitleName: "波斯拉",
		TitleCode: "b_bozra",
	}
}
