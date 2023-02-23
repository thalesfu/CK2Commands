package chilia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥布卢奇察OblucitaBarony struct {
	feud.BaseBarony
}

var BOblucita奥布卢奇察 feud.Barony = &奥布卢奇察OblucitaBarony{}

func init() {
	f := BOblucita奥布卢奇察.(*奥布卢奇察OblucitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oblucita",
		TitleName: "奥布卢奇察",
		TitleCode: "b_oblucita",
	}
}
