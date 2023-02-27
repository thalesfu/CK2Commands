package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德内斯AndenesBarony struct {
	feud.BaseBarony
}

var BAndenes安德内斯 feud.Barony = &安德内斯AndenesBarony{}

func init() {
    f := BAndenes安德内斯.(*安德内斯AndenesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andenes",
		TitleName: "安德内斯",
		TitleCode: "b_andenes",
	}
}
