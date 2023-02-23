package mansura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩提阿利MatiariBarony struct {
	feud.BaseBarony
}

var BMatiari摩提阿利 feud.Barony = &摩提阿利MatiariBarony{}

func init() {
	f := BMatiari摩提阿利.(*摩提阿利MatiariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matiari",
		TitleName: "摩提阿利",
		TitleCode: "b_matiari",
	}
}
