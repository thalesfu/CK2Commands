package gloucester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑尔斯HailesBarony struct {
	feud.BaseBarony
}

var BHailes黑尔斯 feud.Barony = &黑尔斯HailesBarony{}

func init() {
	f := BHailes黑尔斯.(*黑尔斯HailesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hailes",
		TitleName: "黑尔斯",
		TitleCode: "b_hailes",
	}
}
