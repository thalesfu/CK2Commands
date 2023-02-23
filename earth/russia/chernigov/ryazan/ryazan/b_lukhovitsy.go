package ryazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢霍维齐LukhovitsyBarony struct {
	feud.BaseBarony
}

var BLukhovitsy卢霍维齐 feud.Barony = &卢霍维齐LukhovitsyBarony{}

func init() {
	f := BLukhovitsy卢霍维齐.(*卢霍维齐LukhovitsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lukhovitsy",
		TitleName: "卢霍维齐",
		TitleCode: "b_lukhovitsy",
	}
}
