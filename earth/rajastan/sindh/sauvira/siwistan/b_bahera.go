package siwistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆醯罗BaheraBarony struct {
	feud.BaseBarony
}

var BBahera婆醯罗 feud.Barony = &婆醯罗BaheraBarony{}

func init() {
	f := BBahera婆醯罗.(*婆醯罗BaheraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahera",
		TitleName: "婆醯罗",
		TitleCode: "b_bahera",
	}
}
