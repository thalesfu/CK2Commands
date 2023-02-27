package odessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敖德萨OdessaBarony struct {
	feud.BaseBarony
}

var BOdessa敖德萨 feud.Barony = &敖德萨OdessaBarony{}

func init() {
    f := BOdessa敖德萨.(*敖德萨OdessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odessa",
		TitleName: "敖德萨",
		TitleCode: "b_odessa",
	}
}
