package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博德BodoBarony struct {
	feud.BaseBarony
}

var BBodo博德 feud.Barony = &博德BodoBarony{}

func init() {
    f := BBodo博德.(*博德BodoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodo",
		TitleName: "博德",
		TitleCode: "b_bodo",
	}
}
