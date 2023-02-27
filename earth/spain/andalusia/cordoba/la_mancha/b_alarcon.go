package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉扎AlarconBarony struct {
	feud.BaseBarony
}

var BAlarcon阿拉扎 feud.Barony = &阿拉扎AlarconBarony{}

func init() {
    f := BAlarcon阿拉扎.(*阿拉扎AlarconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alarcon",
		TitleName: "阿拉扎",
		TitleCode: "b_alarcon",
	}
}
