package tadla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫扎AfzaBarony struct {
	feud.BaseBarony
}

var BAfza阿夫扎 feud.Barony = &阿夫扎AfzaBarony{}

func init() {
	f := BAfza阿夫扎.(*阿夫扎AfzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "afza",
		TitleName: "阿夫扎",
		TitleCode: "b_afza",
	}
}
