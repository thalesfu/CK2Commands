package mallabhum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦罗吠多GarhbetaBarony struct {
	feud.BaseBarony
}

var BGarhbeta迦罗吠多 feud.Barony = &迦罗吠多GarhbetaBarony{}

func init() {
    f := BGarhbeta迦罗吠多.(*迦罗吠多GarhbetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garhbeta",
		TitleName: "迦罗吠多",
		TitleCode: "b_garhbeta",
	}
}
