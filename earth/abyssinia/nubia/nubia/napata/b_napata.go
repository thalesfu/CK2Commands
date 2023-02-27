package napata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳帕塔NapataBarony struct {
	feud.BaseBarony
}

var BNapata纳帕塔 feud.Barony = &纳帕塔NapataBarony{}

func init() {
    f := BNapata纳帕塔.(*纳帕塔NapataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "napata",
		TitleName: "纳帕塔",
		TitleCode: "b_napata",
	}
}
