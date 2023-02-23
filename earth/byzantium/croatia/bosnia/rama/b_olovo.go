package rama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥洛沃OlovoBarony struct {
	feud.BaseBarony
}

var BOlovo奥洛沃 feud.Barony = &奥洛沃OlovoBarony{}

func init() {
	f := BOlovo奥洛沃.(*奥洛沃OlovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olovo",
		TitleName: "奥洛沃",
		TitleCode: "b_olovo",
	}
}
