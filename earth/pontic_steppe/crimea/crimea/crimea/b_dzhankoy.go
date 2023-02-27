package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 占科伊DzhankoyBarony struct {
	feud.BaseBarony
}

var BDzhankoy占科伊 feud.Barony = &占科伊DzhankoyBarony{}

func init() {
    f := BDzhankoy占科伊.(*占科伊DzhankoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzhankoy",
		TitleName: "占科伊",
		TitleCode: "b_dzhankoy",
	}
}
