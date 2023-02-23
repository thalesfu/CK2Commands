package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉赫里斯TephriceBarony struct {
	feud.BaseBarony
}

var BTephrice特拉赫里斯 feud.Barony = &特拉赫里斯TephriceBarony{}

func init() {
	f := BTephrice特拉赫里斯.(*特拉赫里斯TephriceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tephrice",
		TitleName: "特拉赫里斯",
		TitleCode: "b_tephrice",
	}
}
