package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣贝特朗StbertrandBarony struct {
	feud.BaseBarony
}

var BStbertrand圣贝特朗 feud.Barony = &圣贝特朗StbertrandBarony{}

func init() {
    f := BStbertrand圣贝特朗.(*圣贝特朗StbertrandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stbertrand",
		TitleName: "圣贝特朗",
		TitleCode: "b_stbertrand",
	}
}
