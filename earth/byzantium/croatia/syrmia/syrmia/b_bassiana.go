package syrmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴西亚纳BassianaBarony struct {
	feud.BaseBarony
}

var BBassiana巴西亚纳 feud.Barony = &巴西亚纳BassianaBarony{}

func init() {
    f := BBassiana巴西亚纳.(*巴西亚纳BassianaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bassiana",
		TitleName: "巴西亚纳",
		TitleCode: "b_bassiana",
	}
}
