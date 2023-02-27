package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陶尔米纳TaorminaBarony struct {
	feud.BaseBarony
}

var BTaormina陶尔米纳 feud.Barony = &陶尔米纳TaorminaBarony{}

func init() {
    f := BTaormina陶尔米纳.(*陶尔米纳TaorminaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taormina",
		TitleName: "陶尔米纳",
		TitleCode: "b_taormina",
	}
}
