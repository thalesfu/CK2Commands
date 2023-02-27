package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉孔BraconBarony struct {
	feud.BaseBarony
}

var BBracon布拉孔 feud.Barony = &布拉孔BraconBarony{}

func init() {
    f := BBracon布拉孔.(*布拉孔BraconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bracon",
		TitleName: "布拉孔",
		TitleCode: "b_bracon",
	}
}
