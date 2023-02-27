package balkonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼尔默尔NirmalBarony struct {
	feud.BaseBarony
}

var BNirmal尼尔默尔 feud.Barony = &尼尔默尔NirmalBarony{}

func init() {
    f := BNirmal尼尔默尔.(*尼尔默尔NirmalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nirmal",
		TitleName: "尼尔默尔",
		TitleCode: "b_nirmal",
	}
}
