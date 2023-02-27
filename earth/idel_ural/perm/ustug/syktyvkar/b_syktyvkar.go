package syktyvkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟克特夫卡尔SyktyvkarBarony struct {
	feud.BaseBarony
}

var BSyktyvkar瑟克特夫卡尔 feud.Barony = &瑟克特夫卡尔SyktyvkarBarony{}

func init() {
    f := BSyktyvkar瑟克特夫卡尔.(*瑟克特夫卡尔SyktyvkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syktyvkar",
		TitleName: "瑟克特夫卡尔",
		TitleCode: "b_syktyvkar",
	}
}
