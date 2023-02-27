package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楚格ZugBarony struct {
	feud.BaseBarony
}

var BZug楚格 feud.Barony = &楚格ZugBarony{}

func init() {
    f := BZug楚格.(*楚格ZugBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zug",
		TitleName: "楚格",
		TitleCode: "b_zug",
	}
}
