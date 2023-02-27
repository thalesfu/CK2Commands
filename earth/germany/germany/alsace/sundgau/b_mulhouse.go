package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米卢斯MulhouseBarony struct {
	feud.BaseBarony
}

var BMulhouse米卢斯 feud.Barony = &米卢斯MulhouseBarony{}

func init() {
    f := BMulhouse米卢斯.(*米卢斯MulhouseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mulhouse",
		TitleName: "米卢斯",
		TitleCode: "b_mulhouse",
	}
}
