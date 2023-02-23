package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌摩图尔UmmatturBarony struct {
	feud.BaseBarony
}

var BUmmattur乌摩图尔 feud.Barony = &乌摩图尔UmmatturBarony{}

func init() {
	f := BUmmattur乌摩图尔.(*乌摩图尔UmmatturBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ummattur",
		TitleName: "乌摩图尔",
		TitleCode: "b_ummattur",
	}
}
