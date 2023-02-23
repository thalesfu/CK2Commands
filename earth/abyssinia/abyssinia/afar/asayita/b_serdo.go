package asayita

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟尔多SerdoBarony struct {
	feud.BaseBarony
}

var BSerdo瑟尔多 feud.Barony = &瑟尔多SerdoBarony{}

func init() {
	f := BSerdo瑟尔多.(*瑟尔多SerdoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serdo",
		TitleName: "瑟尔多",
		TitleCode: "b_serdo",
	}
}
