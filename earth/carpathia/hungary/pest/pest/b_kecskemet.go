package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯奇凯梅特KecskemetBarony struct {
	feud.BaseBarony
}

var BKecskemet凯奇凯梅特 feud.Barony = &凯奇凯梅特KecskemetBarony{}

func init() {
    f := BKecskemet凯奇凯梅特.(*凯奇凯梅特KecskemetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kecskemet",
		TitleName: "凯奇凯梅特",
		TitleCode: "b_kecskemet",
	}
}
