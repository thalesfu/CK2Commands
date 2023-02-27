package winchester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安多弗AndoverBarony struct {
	feud.BaseBarony
}

var BAndover安多弗 feud.Barony = &安多弗AndoverBarony{}

func init() {
    f := BAndover安多弗.(*安多弗AndoverBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andover",
		TitleName: "安多弗",
		TitleCode: "b_andover",
	}
}
