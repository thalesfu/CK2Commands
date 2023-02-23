package bath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡德伯里CadburyBarony struct {
	feud.BaseBarony
}

var BCadbury卡德伯里 feud.Barony = &卡德伯里CadburyBarony{}

func init() {
	f := BCadbury卡德伯里.(*卡德伯里CadburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cadbury",
		TitleName: "卡德伯里",
		TitleCode: "b_cadbury",
	}
}
