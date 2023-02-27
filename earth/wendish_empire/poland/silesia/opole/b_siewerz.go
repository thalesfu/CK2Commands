package opole

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢维日SiewerzBarony struct {
	feud.BaseBarony
}

var BSiewerz谢维日 feud.Barony = &谢维日SiewerzBarony{}

func init() {
    f := BSiewerz谢维日.(*谢维日SiewerzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siewerz",
		TitleName: "谢维日",
		TitleCode: "b_siewerz",
	}
}
