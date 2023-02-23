package damoh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佗牟尼DhamoniBarony struct {
	feud.BaseBarony
}

var BDhamoni佗牟尼 feud.Barony = &佗牟尼DhamoniBarony{}

func init() {
	f := BDhamoni佗牟尼.(*佗牟尼DhamoniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhamoni",
		TitleName: "佗牟尼",
		TitleCode: "b_dhamoni",
	}
}
