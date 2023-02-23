package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔ValBarony struct {
	feud.BaseBarony
}

var BVal瓦尔 feud.Barony = &瓦尔ValBarony{}

func init() {
	f := BVal瓦尔.(*瓦尔ValBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "val",
		TitleName: "瓦尔",
		TitleCode: "b_val",
	}
}
