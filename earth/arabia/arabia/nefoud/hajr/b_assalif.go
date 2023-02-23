package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞里夫AssalifBarony struct {
	feud.BaseBarony
}

var BAssalif塞里夫 feud.Barony = &塞里夫AssalifBarony{}

func init() {
	f := BAssalif塞里夫.(*塞里夫AssalifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assalif",
		TitleName: "塞里夫",
		TitleCode: "b_assalif",
	}
}
