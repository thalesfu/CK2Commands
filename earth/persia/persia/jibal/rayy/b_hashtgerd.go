package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈什特盖尔德HashtgerdBarony struct {
	feud.BaseBarony
}

var BHashtgerd哈什特盖尔德 feud.Barony = &哈什特盖尔德HashtgerdBarony{}

func init() {
	f := BHashtgerd哈什特盖尔德.(*哈什特盖尔德HashtgerdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hashtgerd",
		TitleName: "哈什特盖尔德",
		TitleCode: "b_hashtgerd",
	}
}
