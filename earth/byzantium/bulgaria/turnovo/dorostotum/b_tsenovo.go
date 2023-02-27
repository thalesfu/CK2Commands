package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采诺沃TsenovoBarony struct {
	feud.BaseBarony
}

var BTsenovo采诺沃 feud.Barony = &采诺沃TsenovoBarony{}

func init() {
    f := BTsenovo采诺沃.(*采诺沃TsenovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsenovo",
		TitleName: "采诺沃",
		TitleCode: "b_tsenovo",
	}
}
