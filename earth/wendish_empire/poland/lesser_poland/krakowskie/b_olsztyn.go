package krakowskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔什丁OlsztynBarony struct {
	feud.BaseBarony
}

var BOlsztyn奥尔什丁 feud.Barony = &奥尔什丁OlsztynBarony{}

func init() {
    f := BOlsztyn奥尔什丁.(*奥尔什丁OlsztynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olsztyn",
		TitleName: "奥尔什丁",
		TitleCode: "b_olsztyn",
	}
}
