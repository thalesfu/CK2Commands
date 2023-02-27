package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦拉明VaraminBarony struct {
	feud.BaseBarony
}

var BVaramin瓦拉明 feud.Barony = &瓦拉明VaraminBarony{}

func init() {
    f := BVaramin瓦拉明.(*瓦拉明VaraminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varamin",
		TitleName: "瓦拉明",
		TitleCode: "b_varamin",
	}
}
