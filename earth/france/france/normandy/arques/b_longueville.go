package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆格维尔LonguevilleBarony struct {
	feud.BaseBarony
}

var BLongueville隆格维尔 feud.Barony = &隆格维尔LonguevilleBarony{}

func init() {
    f := BLongueville隆格维尔.(*隆格维尔LonguevilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "longueville",
		TitleName: "隆格维尔",
		TitleCode: "b_longueville",
	}
}
