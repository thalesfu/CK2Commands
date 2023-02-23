package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾曼尼亚AljammaniyahBarony struct {
	feud.BaseBarony
}

var BAljammaniyah贾曼尼亚 feud.Barony = &贾曼尼亚AljammaniyahBarony{}

func init() {
	f := BAljammaniyah贾曼尼亚.(*贾曼尼亚AljammaniyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljammaniyah",
		TitleName: "贾曼尼亚",
		TitleCode: "b_aljammaniyah",
	}
}
