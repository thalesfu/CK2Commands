package khijjingakota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔宾基SitabhinjiBarony struct {
	feud.BaseBarony
}

var BSitabhinji斯塔宾基 feud.Barony = &斯塔宾基SitabhinjiBarony{}

func init() {
	f := BSitabhinji斯塔宾基.(*斯塔宾基SitabhinjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sitabhinji",
		TitleName: "斯塔宾基",
		TitleCode: "b_sitabhinji",
	}
}
