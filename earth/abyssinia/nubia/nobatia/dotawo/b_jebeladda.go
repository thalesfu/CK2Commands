package dotawo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰贝腊达JebeladdaBarony struct {
	feud.BaseBarony
}

var BJebeladda杰贝腊达 feud.Barony = &杰贝腊达JebeladdaBarony{}

func init() {
	f := BJebeladda杰贝腊达.(*杰贝腊达JebeladdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jebeladda",
		TitleName: "杰贝腊达",
		TitleCode: "b_jebeladda",
	}
}
