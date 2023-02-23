package moskva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳别尔齐LyubertsyBarony struct {
	feud.BaseBarony
}

var BLyubertsy柳别尔齐 feud.Barony = &柳别尔齐LyubertsyBarony{}

func init() {
	f := BLyubertsy柳别尔齐.(*柳别尔齐LyubertsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyubertsy",
		TitleName: "柳别尔齐",
		TitleCode: "b_lyubertsy",
	}
}
