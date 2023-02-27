package kebbi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉丹阿尔杜Gidan_ardoBarony struct {
	feud.BaseBarony
}

var BGidan_ardo吉丹阿尔杜 feud.Barony = &吉丹阿尔杜Gidan_ardoBarony{}

func init() {
    f := BGidan_ardo吉丹阿尔杜.(*吉丹阿尔杜Gidan_ardoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gidan_ardo",
		TitleName: "吉丹阿尔杜",
		TitleCode: "b_gidan_ardo",
	}
}
