package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳多德拉NadodraBarony struct {
	feud.BaseBarony
}

var BNadodra纳多德拉 feud.Barony = &纳多德拉NadodraBarony{}

func init() {
    f := BNadodra纳多德拉.(*纳多德拉NadodraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nadodra",
		TitleName: "纳多德拉",
		TitleCode: "b_nadodra",
	}
}
