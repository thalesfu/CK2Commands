package kashgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尉头WeitouBarony struct {
	feud.BaseBarony
}

var BWeitou尉头 feud.Barony = &尉头WeitouBarony{}

func init() {
	f := BWeitou尉头.(*尉头WeitouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weitou",
		TitleName: "尉头",
		TitleCode: "b_weitou",
	}
}
