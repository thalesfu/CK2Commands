package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特斯萨拉TessalaBarony struct {
	feud.BaseBarony
}

var BTessala特斯萨拉 feud.Barony = &特斯萨拉TessalaBarony{}

func init() {
    f := BTessala特斯萨拉.(*特斯萨拉TessalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tessala",
		TitleName: "特斯萨拉",
		TitleCode: "b_tessala",
	}
}
