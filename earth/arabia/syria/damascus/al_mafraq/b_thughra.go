package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特胡哈拉ThughraBarony struct {
	feud.BaseBarony
}

var BThughra特胡哈拉 feud.Barony = &特胡哈拉ThughraBarony{}

func init() {
    f := BThughra特胡哈拉.(*特胡哈拉ThughraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thughra",
		TitleName: "特胡哈拉",
		TitleCode: "b_thughra",
	}
}
