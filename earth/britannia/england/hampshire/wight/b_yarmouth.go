package wight

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅茅斯YarmouthBarony struct {
	feud.BaseBarony
}

var BYarmouth雅茅斯 feud.Barony = &雅茅斯YarmouthBarony{}

func init() {
    f := BYarmouth雅茅斯.(*雅茅斯YarmouthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yarmouth",
		TitleName: "雅茅斯",
		TitleCode: "b_yarmouth",
	}
}
