package nice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩纳哥MonacoBarony struct {
	feud.BaseBarony
}

var BMonaco摩纳哥 feud.Barony = &摩纳哥MonacoBarony{}

func init() {
	f := BMonaco摩纳哥.(*摩纳哥MonacoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monaco",
		TitleName: "摩纳哥",
		TitleCode: "b_monaco",
	}
}
