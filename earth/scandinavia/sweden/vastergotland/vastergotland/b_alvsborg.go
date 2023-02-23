package vastergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔夫斯堡AlvsborgBarony struct {
	feud.BaseBarony
}

var BAlvsborg埃尔夫斯堡 feud.Barony = &埃尔夫斯堡AlvsborgBarony{}

func init() {
	f := BAlvsborg埃尔夫斯堡.(*埃尔夫斯堡AlvsborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alvsborg",
		TitleName: "埃尔夫斯堡",
		TitleCode: "b_alvsborg",
	}
}
