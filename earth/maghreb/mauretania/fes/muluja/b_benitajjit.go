package muluja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尼塔吉特BenitajjitBarony struct {
	feud.BaseBarony
}

var BBenitajjit贝尼塔吉特 feud.Barony = &贝尼塔吉特BenitajjitBarony{}

func init() {
	f := BBenitajjit贝尼塔吉特.(*贝尼塔吉特BenitajjitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benitajjit",
		TitleName: "贝尼塔吉特",
		TitleCode: "b_benitajjit",
	}
}
