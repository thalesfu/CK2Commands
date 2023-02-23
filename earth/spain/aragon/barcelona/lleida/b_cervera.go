package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔韦拉CerveraBarony struct {
	feud.BaseBarony
}

var BCervera塞尔韦拉 feud.Barony = &塞尔韦拉CerveraBarony{}

func init() {
	f := BCervera塞尔韦拉.(*塞尔韦拉CerveraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cervera",
		TitleName: "塞尔韦拉",
		TitleCode: "b_cervera",
	}
}
