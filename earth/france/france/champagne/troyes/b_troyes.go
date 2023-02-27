package troyes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特鲁瓦TroyesBarony struct {
	feud.BaseBarony
}

var BTroyes特鲁瓦 feud.Barony = &特鲁瓦TroyesBarony{}

func init() {
    f := BTroyes特鲁瓦.(*特鲁瓦TroyesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "troyes",
		TitleName: "特鲁瓦",
		TitleCode: "b_troyes",
	}
}
