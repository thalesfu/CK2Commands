package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特马罗诺MontemaronoBarony struct {
	feud.BaseBarony
}

var BMontemarono蒙特马罗诺 feud.Barony = &蒙特马罗诺MontemaronoBarony{}

func init() {
	f := BMontemarono蒙特马罗诺.(*蒙特马罗诺MontemaronoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montemarono",
		TitleName: "蒙特马罗诺",
		TitleCode: "b_montemarono",
	}
}
