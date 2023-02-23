package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 高兰陶GalantaBarony struct {
	feud.BaseBarony
}

var BGalanta高兰陶 feud.Barony = &高兰陶GalantaBarony{}

func init() {
	f := BGalanta高兰陶.(*高兰陶GalantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galanta",
		TitleName: "高兰陶",
		TitleCode: "b_galanta",
	}
}
