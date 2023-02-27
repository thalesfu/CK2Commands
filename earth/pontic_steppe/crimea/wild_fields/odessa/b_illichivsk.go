package odessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊利奇夫斯克IllichivskBarony struct {
	feud.BaseBarony
}

var BIllichivsk伊利奇夫斯克 feud.Barony = &伊利奇夫斯克IllichivskBarony{}

func init() {
    f := BIllichivsk伊利奇夫斯克.(*伊利奇夫斯克IllichivskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "illichivsk",
		TitleName: "伊利奇夫斯克",
		TitleCode: "b_illichivsk",
	}
}
