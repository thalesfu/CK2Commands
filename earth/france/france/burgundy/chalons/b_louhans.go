package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢昂LouhansBarony struct {
	feud.BaseBarony
}

var BLouhans卢昂 feud.Barony = &卢昂LouhansBarony{}

func init() {
	f := BLouhans卢昂.(*卢昂LouhansBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "louhans",
		TitleName: "卢昂",
		TitleCode: "b_louhans",
	}
}
