package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖茨黑德GatesheadBarony struct {
	feud.BaseBarony
}

var BGateshead盖茨黑德 feud.Barony = &盖茨黑德GatesheadBarony{}

func init() {
    f := BGateshead盖茨黑德.(*盖茨黑德GatesheadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gateshead",
		TitleName: "盖茨黑德",
		TitleCode: "b_gateshead",
	}
}
