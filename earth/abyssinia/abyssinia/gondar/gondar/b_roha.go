package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗哈RohaBarony struct {
	feud.BaseBarony
}

var BRoha罗哈 feud.Barony = &罗哈RohaBarony{}

func init() {
	f := BRoha罗哈.(*罗哈RohaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roha",
		TitleName: "罗哈",
		TitleCode: "b_roha",
	}
}
