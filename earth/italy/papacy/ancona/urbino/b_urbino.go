package urbino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔比诺UrbinoBarony struct {
	feud.BaseBarony
}

var BUrbino乌尔比诺 feud.Barony = &乌尔比诺UrbinoBarony{}

func init() {
    f := BUrbino乌尔比诺.(*乌尔比诺UrbinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urbino",
		TitleName: "乌尔比诺",
		TitleCode: "b_urbino",
	}
}
