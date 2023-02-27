package kataka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 部婆泥湿伐罗BhubaneswarBarony struct {
	feud.BaseBarony
}

var BBhubaneswar部婆泥湿伐罗 feud.Barony = &部婆泥湿伐罗BhubaneswarBarony{}

func init() {
    f := BBhubaneswar部婆泥湿伐罗.(*部婆泥湿伐罗BhubaneswarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhubaneswar",
		TitleName: "部婆泥湿伐罗",
		TitleCode: "b_bhubaneswar",
	}
}
