package air

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩加尔盖NgargueBarony struct {
	feud.BaseBarony
}

var BNgargue恩加尔盖 feud.Barony = &恩加尔盖NgargueBarony{}

func init() {
	f := BNgargue恩加尔盖.(*恩加尔盖NgargueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ngargue",
		TitleName: "恩加尔盖",
		TitleCode: "b_ngargue",
	}
}
