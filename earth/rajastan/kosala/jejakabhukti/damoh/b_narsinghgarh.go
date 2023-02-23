package damoh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那罗僧诃堡NarsinghgarhBarony struct {
	feud.BaseBarony
}

var BNarsinghgarh那罗僧诃堡 feud.Barony = &那罗僧诃堡NarsinghgarhBarony{}

func init() {
	f := BNarsinghgarh那罗僧诃堡.(*那罗僧诃堡NarsinghgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narsinghgarh",
		TitleName: "那罗僧诃堡",
		TitleCode: "b_narsinghgarh",
	}
}
