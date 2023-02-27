package manyakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲鲁扎巴迪FiruzabadBarony struct {
	feud.BaseBarony
}

var BFiruzabad菲鲁扎巴迪 feud.Barony = &菲鲁扎巴迪FiruzabadBarony{}

func init() {
    f := BFiruzabad菲鲁扎巴迪.(*菲鲁扎巴迪FiruzabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "firuzabad",
		TitleName: "菲鲁扎巴迪",
		TitleCode: "b_firuzabad",
	}
}
