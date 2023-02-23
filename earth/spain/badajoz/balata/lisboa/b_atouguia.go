package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿托吉亚AtouguiaBarony struct {
	feud.BaseBarony
}

var BAtouguia阿托吉亚 feud.Barony = &阿托吉亚AtouguiaBarony{}

func init() {
	f := BAtouguia阿托吉亚.(*阿托吉亚AtouguiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atouguia",
		TitleName: "阿托吉亚",
		TitleCode: "b_atouguia",
	}
}
