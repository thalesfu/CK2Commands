package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗默蒂CromartyBarony struct {
	feud.BaseBarony
}

var BCromarty克罗默蒂 feud.Barony = &克罗默蒂CromartyBarony{}

func init() {
	f := BCromarty克罗默蒂.(*克罗默蒂CromartyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cromarty",
		TitleName: "克罗默蒂",
		TitleCode: "b_cromarty",
	}
}
