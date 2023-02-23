package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达夫尼DaphniBarony struct {
	feud.BaseBarony
}

var BDaphni达夫尼 feud.Barony = &达夫尼DaphniBarony{}

func init() {
	f := BDaphni达夫尼.(*达夫尼DaphniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daphni",
		TitleName: "达夫尼",
		TitleCode: "b_daphni",
	}
}
