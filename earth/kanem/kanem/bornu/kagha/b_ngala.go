package kagha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩加拉NgalaBarony struct {
	feud.BaseBarony
}

var BNgala恩加拉 feud.Barony = &恩加拉NgalaBarony{}

func init() {
    f := BNgala恩加拉.(*恩加拉NgalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ngala",
		TitleName: "恩加拉",
		TitleCode: "b_ngala",
	}
}
