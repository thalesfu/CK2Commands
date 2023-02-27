package pitten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰尔尼茨TernitzBarony struct {
	feud.BaseBarony
}

var BTernitz泰尔尼茨 feud.Barony = &泰尔尼茨TernitzBarony{}

func init() {
    f := BTernitz泰尔尼茨.(*泰尔尼茨TernitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ternitz",
		TitleName: "泰尔尼茨",
		TitleCode: "b_ternitz",
	}
}
