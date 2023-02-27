package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍赫巴ShahbaBarony struct {
	feud.BaseBarony
}

var BShahba舍赫巴 feud.Barony = &舍赫巴ShahbaBarony{}

func init() {
    f := BShahba舍赫巴.(*舍赫巴ShahbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahba",
		TitleName: "舍赫巴",
		TitleCode: "b_shahba",
	}
}
