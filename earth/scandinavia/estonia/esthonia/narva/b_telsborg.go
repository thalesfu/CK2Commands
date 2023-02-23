package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔斯伯格TelsborgBarony struct {
	feud.BaseBarony
}

var BTelsborg特尔斯伯格 feud.Barony = &特尔斯伯格TelsborgBarony{}

func init() {
	f := BTelsborg特尔斯伯格.(*特尔斯伯格TelsborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "telsborg",
		TitleName: "特尔斯伯格",
		TitleCode: "b_telsborg",
	}
}
