package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格利普霍姆GripsholmBarony struct {
	feud.BaseBarony
}

var BGripsholm格利普霍姆 feud.Barony = &格利普霍姆GripsholmBarony{}

func init() {
    f := BGripsholm格利普霍姆.(*格利普霍姆GripsholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gripsholm",
		TitleName: "格利普霍姆",
		TitleCode: "b_gripsholm",
	}
}
