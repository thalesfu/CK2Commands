package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔沃UlvoBarony struct {
	feud.BaseBarony
}

var BUlvo乌尔沃 feud.Barony = &乌尔沃UlvoBarony{}

func init() {
	f := BUlvo乌尔沃.(*乌尔沃UlvoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulvo",
		TitleName: "乌尔沃",
		TitleCode: "b_ulvo",
	}
}
