package semender

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢缅杰尔SemenderBarony struct {
	feud.BaseBarony
}

var BSemender谢缅杰尔 feud.Barony = &谢缅杰尔SemenderBarony{}

func init() {
    f := BSemender谢缅杰尔.(*谢缅杰尔SemenderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semender",
		TitleName: "谢缅杰尔",
		TitleCode: "b_semender",
	}
}
