package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卜泽瓦尔SabzevarBarony struct {
	feud.BaseBarony
}

var BSabzevar萨卜泽瓦尔 feud.Barony = &萨卜泽瓦尔SabzevarBarony{}

func init() {
	f := BSabzevar萨卜泽瓦尔.(*萨卜泽瓦尔SabzevarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabzevar",
		TitleName: "萨卜泽瓦尔",
		TitleCode: "b_sabzevar",
	}
}
