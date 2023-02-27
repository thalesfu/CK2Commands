package connacht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔韦GalwayBarony struct {
	feud.BaseBarony
}

var BGalway戈尔韦 feud.Barony = &戈尔韦GalwayBarony{}

func init() {
    f := BGalway戈尔韦.(*戈尔韦GalwayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galway",
		TitleName: "戈尔韦",
		TitleCode: "b_galway",
	}
}
