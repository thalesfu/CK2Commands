package pereyaslavl_zalessky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫梅利尼基KhmelnikiBarony struct {
	feud.BaseBarony
}

var BKhmelniki赫梅利尼基 feud.Barony = &赫梅利尼基KhmelnikiBarony{}

func init() {
    f := BKhmelniki赫梅利尼基.(*赫梅利尼基KhmelnikiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khmelniki",
		TitleName: "赫梅利尼基",
		TitleCode: "b_khmelniki",
	}
}
