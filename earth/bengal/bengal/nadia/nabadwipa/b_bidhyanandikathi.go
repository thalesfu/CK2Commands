package nabadwipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗陀耶难提迦底BidhyanandikathiBarony struct {
	feud.BaseBarony
}

var BBidhyanandikathi毗陀耶难提迦底 feud.Barony = &毗陀耶难提迦底BidhyanandikathiBarony{}

func init() {
    f := BBidhyanandikathi毗陀耶难提迦底.(*毗陀耶难提迦底BidhyanandikathiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bidhyanandikathi",
		TitleName: "毗陀耶难提迦底",
		TitleCode: "b_bidhyanandikathi",
	}
}
