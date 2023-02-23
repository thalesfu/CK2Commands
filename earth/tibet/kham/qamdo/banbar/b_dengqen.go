package banbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丁青DengqenBarony struct {
	feud.BaseBarony
}

var BDengqen丁青 feud.Barony = &丁青DengqenBarony{}

func init() {
	f := BDengqen丁青.(*丁青DengqenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dengqen",
		TitleName: "丁青",
		TitleCode: "b_dengqen",
	}
}
