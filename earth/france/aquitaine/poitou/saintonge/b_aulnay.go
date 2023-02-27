package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧奈AulnayBarony struct {
	feud.BaseBarony
}

var BAulnay欧奈 feud.Barony = &欧奈AulnayBarony{}

func init() {
    f := BAulnay欧奈.(*欧奈AulnayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aulnay",
		TitleName: "欧奈",
		TitleCode: "b_aulnay",
	}
}
