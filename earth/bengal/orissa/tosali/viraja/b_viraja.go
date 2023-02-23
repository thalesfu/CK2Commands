package viraja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗罗阇VirajaBarony struct {
	feud.BaseBarony
}

var BViraja毗罗阇 feud.Barony = &毗罗阇VirajaBarony{}

func init() {
	f := BViraja毗罗阇.(*毗罗阇VirajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viraja",
		TitleName: "毗罗阇",
		TitleCode: "b_viraja",
	}
}
