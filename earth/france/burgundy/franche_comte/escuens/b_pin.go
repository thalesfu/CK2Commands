package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 潘PinBarony struct {
	feud.BaseBarony
}

var BPin潘 feud.Barony = &潘PinBarony{}

func init() {
    f := BPin潘.(*潘PinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pin",
		TitleName: "潘",
		TitleCode: "b_pin",
	}
}
