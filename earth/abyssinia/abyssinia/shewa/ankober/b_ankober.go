package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安科贝尔AnkoberBarony struct {
	feud.BaseBarony
}

var BAnkober安科贝尔 feud.Barony = &安科贝尔AnkoberBarony{}

func init() {
    f := BAnkober安科贝尔.(*安科贝尔AnkoberBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ankober",
		TitleName: "安科贝尔",
		TitleCode: "b_ankober",
	}
}
