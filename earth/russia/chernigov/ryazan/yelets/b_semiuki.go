package yelets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢米卢基SemiukiBarony struct {
	feud.BaseBarony
}

var BSemiuki谢米卢基 feud.Barony = &谢米卢基SemiukiBarony{}

func init() {
    f := BSemiuki谢米卢基.(*谢米卢基SemiukiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semiuki",
		TitleName: "谢米卢基",
		TitleCode: "b_semiuki",
	}
}
