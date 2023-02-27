package terek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷列克TerekBarony struct {
	feud.BaseBarony
}

var BTerek捷列克 feud.Barony = &捷列克TerekBarony{}

func init() {
    f := BTerek捷列克.(*捷列克TerekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terek",
		TitleName: "捷列克",
		TitleCode: "b_terek",
	}
}
