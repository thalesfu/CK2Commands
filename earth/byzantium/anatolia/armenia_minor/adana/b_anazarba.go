package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿那扎巴AnazarbaBarony struct {
	feud.BaseBarony
}

var BAnazarba阿那扎巴 feud.Barony = &阿那扎巴AnazarbaBarony{}

func init() {
    f := BAnazarba阿那扎巴.(*阿那扎巴AnazarbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anazarba",
		TitleName: "阿那扎巴",
		TitleCode: "b_anazarba",
	}
}
