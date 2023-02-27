package novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺夫哥罗德NovgorodBarony struct {
	feud.BaseBarony
}

var BNovgorod诺夫哥罗德 feud.Barony = &诺夫哥罗德NovgorodBarony{}

func init() {
    f := BNovgorod诺夫哥罗德.(*诺夫哥罗德NovgorodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novgorod",
		TitleName: "诺夫哥罗德",
		TitleCode: "b_novgorod",
	}
}
