package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚奇维日JacwiezBarony struct {
	feud.BaseBarony
}

var BJacwiez亚奇维日 feud.Barony = &亚奇维日JacwiezBarony{}

func init() {
    f := BJacwiez亚奇维日.(*亚奇维日JacwiezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jacwiez",
		TitleName: "亚奇维日",
		TitleCode: "b_jacwiez",
	}
}
