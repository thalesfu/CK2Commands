package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩舍姆EynshamBarony struct {
	feud.BaseBarony
}

var BEynsham恩舍姆 feud.Barony = &恩舍姆EynshamBarony{}

func init() {
	f := BEynsham恩舍姆.(*恩舍姆EynshamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eynsham",
		TitleName: "恩舍姆",
		TitleCode: "b_eynsham",
	}
}
