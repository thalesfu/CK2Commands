package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝拉茨尔克瓦BelacrkvaBarony struct {
	feud.BaseBarony
}

var BBelacrkva贝拉茨尔克瓦 feud.Barony = &贝拉茨尔克瓦BelacrkvaBarony{}

func init() {
    f := BBelacrkva贝拉茨尔克瓦.(*贝拉茨尔克瓦BelacrkvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belacrkva",
		TitleName: "贝拉茨尔克瓦",
		TitleCode: "b_belacrkva",
	}
}
