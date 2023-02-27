package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕耶讷PayerneBarony struct {
	feud.BaseBarony
}

var BPayerne帕耶讷 feud.Barony = &帕耶讷PayerneBarony{}

func init() {
    f := BPayerne帕耶讷.(*帕耶讷PayerneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "payerne",
		TitleName: "帕耶讷",
		TitleCode: "b_payerne",
	}
}
