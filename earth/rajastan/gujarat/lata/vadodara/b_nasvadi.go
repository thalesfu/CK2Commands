package vadodara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那萨迪NasvadiBarony struct {
	feud.BaseBarony
}

var BNasvadi那萨迪 feud.Barony = &那萨迪NasvadiBarony{}

func init() {
    f := BNasvadi那萨迪.(*那萨迪NasvadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nasvadi",
		TitleName: "那萨迪",
		TitleCode: "b_nasvadi",
	}
}
