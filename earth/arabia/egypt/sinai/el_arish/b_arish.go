package el_arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里什ArishBarony struct {
	feud.BaseBarony
}

var BArish阿里什 feud.Barony = &阿里什ArishBarony{}

func init() {
    f := BArish阿里什.(*阿里什ArishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arish",
		TitleName: "阿里什",
		TitleCode: "b_arish",
	}
}
