package eastern_sayan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅热格伊MezhegeyBarony struct {
	feud.BaseBarony
}

var BMezhegey梅热格伊 feud.Barony = &梅热格伊MezhegeyBarony{}

func init() {
    f := BMezhegey梅热格伊.(*梅热格伊MezhegeyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mezhegey",
		TitleName: "梅热格伊",
		TitleCode: "b_mezhegey",
	}
}
