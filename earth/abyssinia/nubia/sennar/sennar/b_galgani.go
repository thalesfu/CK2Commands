package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格甘尼GalganiBarony struct {
	feud.BaseBarony
}

var BGalgani格甘尼 feud.Barony = &格甘尼GalganiBarony{}

func init() {
	f := BGalgani格甘尼.(*格甘尼GalganiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galgani",
		TitleName: "格甘尼",
		TitleCode: "b_galgani",
	}
}
