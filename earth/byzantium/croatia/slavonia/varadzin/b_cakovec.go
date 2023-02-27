package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查科韦茨CakovecBarony struct {
	feud.BaseBarony
}

var BCakovec查科韦茨 feud.Barony = &查科韦茨CakovecBarony{}

func init() {
    f := BCakovec查科韦茨.(*查科韦茨CakovecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cakovec",
		TitleName: "查科韦茨",
		TitleCode: "b_cakovec",
	}
}
