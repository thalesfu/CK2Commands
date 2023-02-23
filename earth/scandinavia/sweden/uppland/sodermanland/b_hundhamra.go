package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洪德哈姆拉HundhamraBarony struct {
	feud.BaseBarony
}

var BHundhamra洪德哈姆拉 feud.Barony = &洪德哈姆拉HundhamraBarony{}

func init() {
	f := BHundhamra洪德哈姆拉.(*洪德哈姆拉HundhamraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hundhamra",
		TitleName: "洪德哈姆拉",
		TitleCode: "b_hundhamra",
	}
}
