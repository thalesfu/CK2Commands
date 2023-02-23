package tregor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普卢贡旺PlougonvenBarony struct {
	feud.BaseBarony
}

var BPlougonven普卢贡旺 feud.Barony = &普卢贡旺PlougonvenBarony{}

func init() {
	f := BPlougonven普卢贡旺.(*普卢贡旺PlougonvenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plougonven",
		TitleName: "普卢贡旺",
		TitleCode: "b_plougonven",
	}
}
