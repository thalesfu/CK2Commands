package nevers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉沙里泰LachariteBarony struct {
	feud.BaseBarony
}

var BLacharite拉沙里泰 feud.Barony = &拉沙里泰LachariteBarony{}

func init() {
	f := BLacharite拉沙里泰.(*拉沙里泰LachariteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lacharite",
		TitleName: "拉沙里泰",
		TitleCode: "b_lacharite",
	}
}
