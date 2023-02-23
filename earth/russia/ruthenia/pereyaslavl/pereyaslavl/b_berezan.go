package pereyaslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列赞BerezanBarony struct {
	feud.BaseBarony
}

var BBerezan别列赞 feud.Barony = &别列赞BerezanBarony{}

func init() {
	f := BBerezan别列赞.(*别列赞BerezanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berezan",
		TitleName: "别列赞",
		TitleCode: "b_berezan",
	}
}
