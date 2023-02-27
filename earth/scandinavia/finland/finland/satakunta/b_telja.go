package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰列TeljaBarony struct {
	feud.BaseBarony
}

var BTelja泰列 feud.Barony = &泰列TeljaBarony{}

func init() {
    f := BTelja泰列.(*泰列TeljaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "telja",
		TitleName: "泰列",
		TitleCode: "b_telja",
	}
}
