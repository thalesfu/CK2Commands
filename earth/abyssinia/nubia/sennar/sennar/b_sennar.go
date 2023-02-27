package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 森纳尔SennarBarony struct {
	feud.BaseBarony
}

var BSennar森纳尔 feud.Barony = &森纳尔SennarBarony{}

func init() {
    f := BSennar森纳尔.(*森纳尔SennarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sennar",
		TitleName: "森纳尔",
		TitleCode: "b_sennar",
	}
}
