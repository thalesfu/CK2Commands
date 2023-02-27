package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯法克斯SfaxBarony struct {
	feud.BaseBarony
}

var BSfax斯法克斯 feud.Barony = &斯法克斯SfaxBarony{}

func init() {
    f := BSfax斯法克斯.(*斯法克斯SfaxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sfax",
		TitleName: "斯法克斯",
		TitleCode: "b_sfax",
	}
}
