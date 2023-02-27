package tobol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔拉克CherlakBarony struct {
	feud.BaseBarony
}

var BCherlak切尔拉克 feud.Barony = &切尔拉克CherlakBarony{}

func init() {
    f := BCherlak切尔拉克.(*切尔拉克CherlakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherlak",
		TitleName: "切尔拉克",
		TitleCode: "b_cherlak",
	}
}
