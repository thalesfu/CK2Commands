package targu_jiu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴亚德阿拉默Baia_de_aramaBarony struct {
	feud.BaseBarony
}

var BBaia_de_arama巴亚德阿拉默 feud.Barony = &巴亚德阿拉默Baia_de_aramaBarony{}

func init() {
    f := BBaia_de_arama巴亚德阿拉默.(*巴亚德阿拉默Baia_de_aramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baia_de_arama",
		TitleName: "巴亚德阿拉默",
		TitleCode: "b_baia_de_arama",
	}
}
