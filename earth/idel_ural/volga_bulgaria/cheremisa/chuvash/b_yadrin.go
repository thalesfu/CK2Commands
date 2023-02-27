package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚德林YadrinBarony struct {
	feud.BaseBarony
}

var BYadrin亚德林 feud.Barony = &亚德林YadrinBarony{}

func init() {
    f := BYadrin亚德林.(*亚德林YadrinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yadrin",
		TitleName: "亚德林",
		TitleCode: "b_yadrin",
	}
}
