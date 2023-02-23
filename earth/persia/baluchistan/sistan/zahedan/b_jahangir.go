package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾汉吉尔JahangirBarony struct {
	feud.BaseBarony
}

var BJahangir贾汉吉尔 feud.Barony = &贾汉吉尔JahangirBarony{}

func init() {
	f := BJahangir贾汉吉尔.(*贾汉吉尔JahangirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jahangir",
		TitleName: "贾汉吉尔",
		TitleCode: "b_jahangir",
	}
}
