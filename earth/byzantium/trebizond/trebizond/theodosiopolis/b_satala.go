package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨塔拉SatalaBarony struct {
	feud.BaseBarony
}

var BSatala萨塔拉 feud.Barony = &萨塔拉SatalaBarony{}

func init() {
    f := BSatala萨塔拉.(*萨塔拉SatalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satala",
		TitleName: "萨塔拉",
		TitleCode: "b_satala",
	}
}
