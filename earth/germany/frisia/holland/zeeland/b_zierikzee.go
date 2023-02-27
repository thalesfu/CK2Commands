package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济里克泽ZierikzeeBarony struct {
	feud.BaseBarony
}

var BZierikzee济里克泽 feud.Barony = &济里克泽ZierikzeeBarony{}

func init() {
    f := BZierikzee济里克泽.(*济里克泽ZierikzeeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zierikzee",
		TitleName: "济里克泽",
		TitleCode: "b_zierikzee",
	}
}
