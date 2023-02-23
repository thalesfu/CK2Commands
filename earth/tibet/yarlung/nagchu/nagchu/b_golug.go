package nagchu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古露GolugBarony struct {
	feud.BaseBarony
}

var BGolug古露 feud.Barony = &古露GolugBarony{}

func init() {
	f := BGolug古露.(*古露GolugBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "golug",
		TitleName: "古露",
		TitleCode: "b_golug",
	}
}
