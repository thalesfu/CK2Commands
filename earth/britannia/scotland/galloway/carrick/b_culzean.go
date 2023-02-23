package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克伦CulzeanBarony struct {
	feud.BaseBarony
}

var BCulzean克伦 feud.Barony = &克伦CulzeanBarony{}

func init() {
	f := BCulzean克伦.(*克伦CulzeanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "culzean",
		TitleName: "克伦",
		TitleCode: "b_culzean",
	}
}
