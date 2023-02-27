package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特卢克TeluchBarony struct {
	feud.BaseBarony
}

var BTeluch特卢克 feud.Barony = &特卢克TeluchBarony{}

func init() {
    f := BTeluch特卢克.(*特卢克TeluchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teluch",
		TitleName: "特卢克",
		TitleCode: "b_teluch",
	}
}
