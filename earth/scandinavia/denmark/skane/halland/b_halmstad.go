package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔姆斯塔德HalmstadBarony struct {
	feud.BaseBarony
}

var BHalmstad哈尔姆斯塔德 feud.Barony = &哈尔姆斯塔德HalmstadBarony{}

func init() {
    f := BHalmstad哈尔姆斯塔德.(*哈尔姆斯塔德HalmstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halmstad",
		TitleName: "哈尔姆斯塔德",
		TitleCode: "b_halmstad",
	}
}
