package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赞迪QundiBarony struct {
	feud.BaseBarony
}

var BQundi赞迪 feud.Barony = &赞迪QundiBarony{}

func init() {
    f := BQundi赞迪.(*赞迪QundiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qundi",
		TitleName: "赞迪",
		TitleCode: "b_qundi",
	}
}
