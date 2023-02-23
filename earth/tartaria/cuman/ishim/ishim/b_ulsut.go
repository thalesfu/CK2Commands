package ishim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔苏特UlsutBarony struct {
	feud.BaseBarony
}

var BUlsut乌尔苏特 feud.Barony = &乌尔苏特UlsutBarony{}

func init() {
	f := BUlsut乌尔苏特.(*乌尔苏特UlsutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulsut",
		TitleName: "乌尔苏特",
		TitleCode: "b_ulsut",
	}
}
