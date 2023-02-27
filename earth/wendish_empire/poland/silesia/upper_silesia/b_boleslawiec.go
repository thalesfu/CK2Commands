package upper_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博莱斯瓦维茨BoleslawiecBarony struct {
	feud.BaseBarony
}

var BBoleslawiec博莱斯瓦维茨 feud.Barony = &博莱斯瓦维茨BoleslawiecBarony{}

func init() {
    f := BBoleslawiec博莱斯瓦维茨.(*博莱斯瓦维茨BoleslawiecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boleslawiec",
		TitleName: "博莱斯瓦维茨",
		TitleCode: "b_boleslawiec",
	}
}
