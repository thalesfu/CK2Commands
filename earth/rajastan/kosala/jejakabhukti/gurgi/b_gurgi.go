package gurgi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古卢耆GurgiBarony struct {
	feud.BaseBarony
}

var BGurgi古卢耆 feud.Barony = &古卢耆GurgiBarony{}

func init() {
    f := BGurgi古卢耆.(*古卢耆GurgiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurgi",
		TitleName: "古卢耆",
		TitleCode: "b_gurgi",
	}
}
