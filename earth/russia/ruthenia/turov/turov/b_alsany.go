package turov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥利沙内AlsanyBarony struct {
	feud.BaseBarony
}

var BAlsany奥利沙内 feud.Barony = &奥利沙内AlsanyBarony{}

func init() {
	f := BAlsany奥利沙内.(*奥利沙内AlsanyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alsany",
		TitleName: "奥利沙内",
		TitleCode: "b_alsany",
	}
}
