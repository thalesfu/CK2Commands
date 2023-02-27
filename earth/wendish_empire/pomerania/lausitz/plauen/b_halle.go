package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈雷HalleBarony struct {
	feud.BaseBarony
}

var BHalle哈雷 feud.Barony = &哈雷HalleBarony{}

func init() {
    f := BHalle哈雷.(*哈雷HalleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halle",
		TitleName: "哈雷",
		TitleCode: "b_halle",
	}
}
