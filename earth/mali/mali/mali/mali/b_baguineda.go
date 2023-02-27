package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴吉奈达BaguinedaBarony struct {
	feud.BaseBarony
}

var BBaguineda巴吉奈达 feud.Barony = &巴吉奈达BaguinedaBarony{}

func init() {
    f := BBaguineda巴吉奈达.(*巴吉奈达BaguinedaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baguineda",
		TitleName: "巴吉奈达",
		TitleCode: "b_baguineda",
	}
}
