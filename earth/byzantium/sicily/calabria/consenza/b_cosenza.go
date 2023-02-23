package consenza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科森扎CosenzaBarony struct {
	feud.BaseBarony
}

var BCosenza科森扎 feud.Barony = &科森扎CosenzaBarony{}

func init() {
	f := BCosenza科森扎.(*科森扎CosenzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cosenza",
		TitleName: "科森扎",
		TitleCode: "b_cosenza",
	}
}
