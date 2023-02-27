package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马罗米察MaromitsaBarony struct {
	feud.BaseBarony
}

var BMaromitsa马罗米察 feud.Barony = &马罗米察MaromitsaBarony{}

func init() {
    f := BMaromitsa马罗米察.(*马罗米察MaromitsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maromitsa",
		TitleName: "马罗米察",
		TitleCode: "b_maromitsa",
	}
}
