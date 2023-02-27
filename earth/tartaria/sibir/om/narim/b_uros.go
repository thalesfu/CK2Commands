package narim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌罗斯UrosBarony struct {
	feud.BaseBarony
}

var BUros乌罗斯 feud.Barony = &乌罗斯UrosBarony{}

func init() {
    f := BUros乌罗斯.(*乌罗斯UrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uros",
		TitleName: "乌罗斯",
		TitleCode: "b_uros",
	}
}
