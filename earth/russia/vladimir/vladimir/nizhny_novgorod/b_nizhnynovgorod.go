package nizhny_novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下诺夫哥罗德NizhnynovgorodBarony struct {
	feud.BaseBarony
}

var BNizhnynovgorod下诺夫哥罗德 feud.Barony = &下诺夫哥罗德NizhnynovgorodBarony{}

func init() {
    f := BNizhnynovgorod下诺夫哥罗德.(*下诺夫哥罗德NizhnynovgorodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizhnynovgorod",
		TitleName: "下诺夫哥罗德",
		TitleCode: "b_nizhnynovgorod",
	}
}
