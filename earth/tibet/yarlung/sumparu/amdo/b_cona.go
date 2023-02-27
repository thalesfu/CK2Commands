package amdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 错那ConaBarony struct {
	feud.BaseBarony
}

var BCona错那 feud.Barony = &错那ConaBarony{}

func init() {
    f := BCona错那.(*错那ConaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cona",
		TitleName: "错那",
		TitleCode: "b_cona",
	}
}
