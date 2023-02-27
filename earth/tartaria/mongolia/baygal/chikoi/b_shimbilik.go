package chikoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 申比利克ShimbilikBarony struct {
	feud.BaseBarony
}

var BShimbilik申比利克 feud.Barony = &申比利克ShimbilikBarony{}

func init() {
    f := BShimbilik申比利克.(*申比利克ShimbilikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shimbilik",
		TitleName: "申比利克",
		TitleCode: "b_shimbilik",
	}
}
