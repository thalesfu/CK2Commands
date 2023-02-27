package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔扎兹KerzazBarony struct {
	feud.BaseBarony
}

var BKerzaz克尔扎兹 feud.Barony = &克尔扎兹KerzazBarony{}

func init() {
    f := BKerzaz克尔扎兹.(*克尔扎兹KerzazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerzaz",
		TitleName: "克尔扎兹",
		TitleCode: "b_kerzaz",
	}
}
