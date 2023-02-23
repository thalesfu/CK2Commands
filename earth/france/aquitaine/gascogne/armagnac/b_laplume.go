package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉普吕姆LaplumeBarony struct {
	feud.BaseBarony
}

var BLaplume拉普吕姆 feud.Barony = &拉普吕姆LaplumeBarony{}

func init() {
	f := BLaplume拉普吕姆.(*拉普吕姆LaplumeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laplume",
		TitleName: "拉普吕姆",
		TitleCode: "b_laplume",
	}
}
