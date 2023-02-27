package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利姆LympneBarony struct {
	feud.BaseBarony
}

var BLympne利姆 feud.Barony = &利姆LympneBarony{}

func init() {
    f := BLympne利姆.(*利姆LympneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lympne",
		TitleName: "利姆",
		TitleCode: "b_lympne",
	}
}
