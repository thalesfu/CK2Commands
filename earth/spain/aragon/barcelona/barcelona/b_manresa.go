package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼雷萨ManresaBarony struct {
	feud.BaseBarony
}

var BManresa曼雷萨 feud.Barony = &曼雷萨ManresaBarony{}

func init() {
	f := BManresa曼雷萨.(*曼雷萨ManresaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manresa",
		TitleName: "曼雷萨",
		TitleCode: "b_manresa",
	}
}
