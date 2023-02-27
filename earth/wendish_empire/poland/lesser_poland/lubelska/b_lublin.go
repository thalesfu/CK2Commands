package lubelska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢布林LublinBarony struct {
	feud.BaseBarony
}

var BLublin卢布林 feud.Barony = &卢布林LublinBarony{}

func init() {
    f := BLublin卢布林.(*卢布林LublinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lublin",
		TitleName: "卢布林",
		TitleCode: "b_lublin",
	}
}
