package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣马尔科达伦齐奥SanmarcodalunzioBarony struct {
	feud.BaseBarony
}

var BSanmarcodalunzio圣马尔科达伦齐奥 feud.Barony = &圣马尔科达伦齐奥SanmarcodalunzioBarony{}

func init() {
    f := BSanmarcodalunzio圣马尔科达伦齐奥.(*圣马尔科达伦齐奥SanmarcodalunzioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanmarcodalunzio",
		TitleName: "圣马尔科达伦齐奥",
		TitleCode: "b_sanmarcodalunzio",
	}
}
