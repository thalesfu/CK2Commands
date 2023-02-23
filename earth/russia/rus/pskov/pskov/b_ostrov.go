package pskov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特罗夫OstrovBarony struct {
	feud.BaseBarony
}

var BOstrov奥斯特罗夫 feud.Barony = &奥斯特罗夫OstrovBarony{}

func init() {
	f := BOstrov奥斯特罗夫.(*奥斯特罗夫OstrovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostrov",
		TitleName: "奥斯特罗夫",
		TitleCode: "b_ostrov",
	}
}
