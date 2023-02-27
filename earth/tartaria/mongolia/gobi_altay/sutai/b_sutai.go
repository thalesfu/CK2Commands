package sutai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏泰SutaiBarony struct {
	feud.BaseBarony
}

var BSutai苏泰 feud.Barony = &苏泰SutaiBarony{}

func init() {
    f := BSutai苏泰.(*苏泰SutaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sutai",
		TitleName: "苏泰",
		TitleCode: "b_sutai",
	}
}
