package buzachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克赞库尔Kyzan_kulBarony struct {
	feud.BaseBarony
}

var BKyzan_kul克赞库尔 feud.Barony = &克赞库尔Kyzan_kulBarony{}

func init() {
    f := BKyzan_kul克赞库尔.(*克赞库尔Kyzan_kulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzan_kul",
		TitleName: "克赞库尔",
		TitleCode: "b_kyzan_kul",
	}
}
