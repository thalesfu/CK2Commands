package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安特卫普AntwerpenBarony struct {
	feud.BaseBarony
}

var BAntwerpen安特卫普 feud.Barony = &安特卫普AntwerpenBarony{}

func init() {
    f := BAntwerpen安特卫普.(*安特卫普AntwerpenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antwerpen",
		TitleName: "安特卫普",
		TitleCode: "b_antwerpen",
	}
}
