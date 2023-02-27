package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔斯塔德BolstadBarony struct {
	feud.BaseBarony
}

var BBolstad博尔斯塔德 feud.Barony = &博尔斯塔德BolstadBarony{}

func init() {
    f := BBolstad博尔斯塔德.(*博尔斯塔德BolstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolstad",
		TitleName: "博尔斯塔德",
		TitleCode: "b_bolstad",
	}
}
