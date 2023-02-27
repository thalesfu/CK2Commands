package rostock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔欣MalchinBarony struct {
	feud.BaseBarony
}

var BMalchin马尔欣 feud.Barony = &马尔欣MalchinBarony{}

func init() {
    f := BMalchin马尔欣.(*马尔欣MalchinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malchin",
		TitleName: "马尔欣",
		TitleCode: "b_malchin",
	}
}
