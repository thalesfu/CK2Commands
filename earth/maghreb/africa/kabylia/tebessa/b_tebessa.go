package tebessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰贝萨TebessaBarony struct {
	feud.BaseBarony
}

var BTebessa泰贝萨 feud.Barony = &泰贝萨TebessaBarony{}

func init() {
    f := BTebessa泰贝萨.(*泰贝萨TebessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tebessa",
		TitleName: "泰贝萨",
		TitleCode: "b_tebessa",
	}
}
