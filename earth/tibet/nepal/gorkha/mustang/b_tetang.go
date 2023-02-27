package mustang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谛当TetangBarony struct {
	feud.BaseBarony
}

var BTetang谛当 feud.Barony = &谛当TetangBarony{}

func init() {
    f := BTetang谛当.(*谛当TetangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tetang",
		TitleName: "谛当",
		TitleCode: "b_tetang",
	}
}
