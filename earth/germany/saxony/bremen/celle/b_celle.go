package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 策勒CelleBarony struct {
	feud.BaseBarony
}

var BCelle策勒 feud.Barony = &策勒CelleBarony{}

func init() {
    f := BCelle策勒.(*策勒CelleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "celle",
		TitleName: "策勒",
		TitleCode: "b_celle",
	}
}
