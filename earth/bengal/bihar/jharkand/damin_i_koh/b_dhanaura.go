package damin_i_koh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀奴罗DhanauraBarony struct {
	feud.BaseBarony
}

var BDhanaura陀奴罗 feud.Barony = &陀奴罗DhanauraBarony{}

func init() {
    f := BDhanaura陀奴罗.(*陀奴罗DhanauraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhanaura",
		TitleName: "陀奴罗",
		TitleCode: "b_dhanaura",
	}
}
