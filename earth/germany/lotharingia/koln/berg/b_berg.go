package berg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝格BergBarony struct {
	feud.BaseBarony
}

var BBerg贝格 feud.Barony = &贝格BergBarony{}

func init() {
	f := BBerg贝格.(*贝格BergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berg",
		TitleName: "贝格",
		TitleCode: "b_berg",
	}
}
