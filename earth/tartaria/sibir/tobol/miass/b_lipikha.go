package miass

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利皮哈LipikhaBarony struct {
	feud.BaseBarony
}

var BLipikha利皮哈 feud.Barony = &利皮哈LipikhaBarony{}

func init() {
    f := BLipikha利皮哈.(*利皮哈LipikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lipikha",
		TitleName: "利皮哈",
		TitleCode: "b_lipikha",
	}
}
