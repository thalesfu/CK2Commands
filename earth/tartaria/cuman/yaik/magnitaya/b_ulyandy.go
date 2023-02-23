package magnitaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌良德UlyandyBarony struct {
	feud.BaseBarony
}

var BUlyandy乌良德 feud.Barony = &乌良德UlyandyBarony{}

func init() {
	f := BUlyandy乌良德.(*乌良德UlyandyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulyandy",
		TitleName: "乌良德",
		TitleCode: "b_ulyandy",
	}
}
