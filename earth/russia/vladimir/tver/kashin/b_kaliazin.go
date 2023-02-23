package kashin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利亚津KaliazinBarony struct {
	feud.BaseBarony
}

var BKaliazin卡利亚津 feud.Barony = &卡利亚津KaliazinBarony{}

func init() {
	f := BKaliazin卡利亚津.(*卡利亚津KaliazinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaliazin",
		TitleName: "卡利亚津",
		TitleCode: "b_kaliazin",
	}
}
