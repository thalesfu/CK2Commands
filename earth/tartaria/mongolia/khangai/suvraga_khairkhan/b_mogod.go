package suvraga_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛高德MogodBarony struct {
	feud.BaseBarony
}

var BMogod毛高德 feud.Barony = &毛高德MogodBarony{}

func init() {
    f := BMogod毛高德.(*毛高德MogodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mogod",
		TitleName: "毛高德",
		TitleCode: "b_mogod",
	}
}
