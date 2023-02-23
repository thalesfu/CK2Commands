package dunois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕泰PatayBarony struct {
	feud.BaseBarony
}

var BPatay帕泰 feud.Barony = &帕泰PatayBarony{}

func init() {
	f := BPatay帕泰.(*帕泰PatayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patay",
		TitleName: "帕泰",
		TitleCode: "b_patay",
	}
}
