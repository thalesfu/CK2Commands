package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克隆马克诺伊斯ClonmacnoiseBarony struct {
	feud.BaseBarony
}

var BClonmacnoise克隆马克诺伊斯 feud.Barony = &克隆马克诺伊斯ClonmacnoiseBarony{}

func init() {
	f := BClonmacnoise克隆马克诺伊斯.(*克隆马克诺伊斯ClonmacnoiseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clonmacnoise",
		TitleName: "克隆马克诺伊斯",
		TitleCode: "b_clonmacnoise",
	}
}
