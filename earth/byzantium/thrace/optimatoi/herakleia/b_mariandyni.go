package herakleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马里安底尼MariandyniBarony struct {
	feud.BaseBarony
}

var BMariandyni马里安底尼 feud.Barony = &马里安底尼MariandyniBarony{}

func init() {
	f := BMariandyni马里安底尼.(*马里安底尼MariandyniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mariandyni",
		TitleName: "马里安底尼",
		TitleCode: "b_mariandyni",
	}
}
