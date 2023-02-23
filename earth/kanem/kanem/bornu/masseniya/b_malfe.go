package masseniya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马勒非MalfeBarony struct {
	feud.BaseBarony
}

var BMalfe马勒非 feud.Barony = &马勒非MalfeBarony{}

func init() {
	f := BMalfe马勒非.(*马勒非MalfeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malfe",
		TitleName: "马勒非",
		TitleCode: "b_malfe",
	}
}
