package haruppeswara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查莱碉CharaideoBarony struct {
	feud.BaseBarony
}

var BCharaideo查莱碉 feud.Barony = &查莱碉CharaideoBarony{}

func init() {
	f := BCharaideo查莱碉.(*查莱碉CharaideoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charaideo",
		TitleName: "查莱碉",
		TitleCode: "b_charaideo",
	}
}
