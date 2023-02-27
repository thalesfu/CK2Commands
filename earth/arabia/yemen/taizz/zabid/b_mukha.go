package zabid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆哈MukhaBarony struct {
	feud.BaseBarony
}

var BMukha穆哈 feud.Barony = &穆哈MukhaBarony{}

func init() {
    f := BMukha穆哈.(*穆哈MukhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mukha",
		TitleName: "穆哈",
		TitleCode: "b_mukha",
	}
}
