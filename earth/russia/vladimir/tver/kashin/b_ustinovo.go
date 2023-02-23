package kashin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季诺沃UstinovoBarony struct {
	feud.BaseBarony
}

var BUstinovo乌斯季诺沃 feud.Barony = &乌斯季诺沃UstinovoBarony{}

func init() {
	f := BUstinovo乌斯季诺沃.(*乌斯季诺沃UstinovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustinovo",
		TitleName: "乌斯季诺沃",
		TitleCode: "b_ustinovo",
	}
}
