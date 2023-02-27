package aoudaghost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔卡堡KsarelbarkaBarony struct {
	feud.BaseBarony
}

var BKsarelbarka巴尔卡堡 feud.Barony = &巴尔卡堡KsarelbarkaBarony{}

func init() {
    f := BKsarelbarka巴尔卡堡.(*巴尔卡堡KsarelbarkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ksarelbarka",
		TitleName: "巴尔卡堡",
		TitleCode: "b_ksarelbarka",
	}
}
