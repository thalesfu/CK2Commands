package goalpara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦摩佉耶KamakhyaBarony struct {
	feud.BaseBarony
}

var BKamakhya迦摩佉耶 feud.Barony = &迦摩佉耶KamakhyaBarony{}

func init() {
    f := BKamakhya迦摩佉耶.(*迦摩佉耶KamakhyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamakhya",
		TitleName: "迦摩佉耶",
		TitleCode: "b_kamakhya",
	}
}
