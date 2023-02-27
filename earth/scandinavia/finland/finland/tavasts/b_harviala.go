package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔维亚拉HarvialaBarony struct {
	feud.BaseBarony
}

var BHarviala哈尔维亚拉 feud.Barony = &哈尔维亚拉HarvialaBarony{}

func init() {
    f := BHarviala哈尔维亚拉.(*哈尔维亚拉HarvialaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harviala",
		TitleName: "哈尔维亚拉",
		TitleCode: "b_harviala",
	}
}
