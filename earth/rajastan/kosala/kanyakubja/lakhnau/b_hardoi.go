package lakhnau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃罗豆伊HardoiBarony struct {
	feud.BaseBarony
}

var BHardoi诃罗豆伊 feud.Barony = &诃罗豆伊HardoiBarony{}

func init() {
    f := BHardoi诃罗豆伊.(*诃罗豆伊HardoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hardoi",
		TitleName: "诃罗豆伊",
		TitleCode: "b_hardoi",
	}
}
