package shish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姆索诺沃SamsonovoBarony struct {
	feud.BaseBarony
}

var BSamsonovo萨姆索诺沃 feud.Barony = &萨姆索诺沃SamsonovoBarony{}

func init() {
    f := BSamsonovo萨姆索诺沃.(*萨姆索诺沃SamsonovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samsonovo",
		TitleName: "萨姆索诺沃",
		TitleCode: "b_samsonovo",
	}
}
