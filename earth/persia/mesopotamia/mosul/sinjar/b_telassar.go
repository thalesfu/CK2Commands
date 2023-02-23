package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉萨TelassarBarony struct {
	feud.BaseBarony
}

var BTelassar特拉萨 feud.Barony = &特拉萨TelassarBarony{}

func init() {
	f := BTelassar特拉萨.(*特拉萨TelassarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "telassar",
		TitleName: "特拉萨",
		TitleCode: "b_telassar",
	}
}
