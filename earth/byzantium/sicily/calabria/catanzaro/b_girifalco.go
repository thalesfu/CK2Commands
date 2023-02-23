package catanzaro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉里法尔科GirifalcoBarony struct {
	feud.BaseBarony
}

var BGirifalco吉里法尔科 feud.Barony = &吉里法尔科GirifalcoBarony{}

func init() {
	f := BGirifalco吉里法尔科.(*吉里法尔科GirifalcoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "girifalco",
		TitleName: "吉里法尔科",
		TitleCode: "b_girifalco",
	}
}
