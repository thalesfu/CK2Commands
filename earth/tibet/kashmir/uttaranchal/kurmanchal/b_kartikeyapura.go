package kurmanchal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦絺吉夜补罗KartikeyapuraBarony struct {
	feud.BaseBarony
}

var BKartikeyapura迦絺吉夜补罗 feud.Barony = &迦絺吉夜补罗KartikeyapuraBarony{}

func init() {
	f := BKartikeyapura迦絺吉夜补罗.(*迦絺吉夜补罗KartikeyapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kartikeyapura",
		TitleName: "迦絺吉夜补罗",
		TitleCode: "b_kartikeyapura",
	}
}
