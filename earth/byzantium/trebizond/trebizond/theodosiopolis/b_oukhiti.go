package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夸西提OukhitiBarony struct {
	feud.BaseBarony
}

var BOukhiti夸西提 feud.Barony = &夸西提OukhitiBarony{}

func init() {
    f := BOukhiti夸西提.(*夸西提OukhitiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oukhiti",
		TitleName: "夸西提",
		TitleCode: "b_oukhiti",
	}
}
