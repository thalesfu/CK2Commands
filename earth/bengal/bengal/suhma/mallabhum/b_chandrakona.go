package mallabhum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀罗拘拿ChandrakonaBarony struct {
	feud.BaseBarony
}

var BChandrakona旃陀罗拘拿 feud.Barony = &旃陀罗拘拿ChandrakonaBarony{}

func init() {
    f := BChandrakona旃陀罗拘拿.(*旃陀罗拘拿ChandrakonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chandrakona",
		TitleName: "旃陀罗拘拿",
		TitleCode: "b_chandrakona",
	}
}
