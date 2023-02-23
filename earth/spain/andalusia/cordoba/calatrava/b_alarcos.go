package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉尔科斯AlarcosBarony struct {
	feud.BaseBarony
}

var BAlarcos阿拉尔科斯 feud.Barony = &阿拉尔科斯AlarcosBarony{}

func init() {
	f := BAlarcos阿拉尔科斯.(*阿拉尔科斯AlarcosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alarcos",
		TitleName: "阿拉尔科斯",
		TitleCode: "b_alarcos",
	}
}
