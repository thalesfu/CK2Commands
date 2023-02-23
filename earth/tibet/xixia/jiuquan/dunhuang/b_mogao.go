package dunhuang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫高MogaoBarony struct {
	feud.BaseBarony
}

var BMogao莫高 feud.Barony = &莫高MogaoBarony{}

func init() {
	f := BMogao莫高.(*莫高MogaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mogao",
		TitleName: "莫高",
		TitleCode: "b_mogao",
	}
}
