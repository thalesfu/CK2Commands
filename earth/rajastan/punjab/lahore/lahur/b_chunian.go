package lahur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楚里安ChunianBarony struct {
	feud.BaseBarony
}

var BChunian楚里安 feud.Barony = &楚里安ChunianBarony{}

func init() {
	f := BChunian楚里安.(*楚里安ChunianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chunian",
		TitleName: "楚里安",
		TitleCode: "b_chunian",
	}
}
