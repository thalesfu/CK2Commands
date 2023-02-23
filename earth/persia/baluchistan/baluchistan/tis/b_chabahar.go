package tis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈赫巴哈尔ChabaharBarony struct {
	feud.BaseBarony
}

var BChabahar哈赫巴哈尔 feud.Barony = &哈赫巴哈尔ChabaharBarony{}

func init() {
	f := BChabahar哈赫巴哈尔.(*哈赫巴哈尔ChabaharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chabahar",
		TitleName: "哈赫巴哈尔",
		TitleCode: "b_chabahar",
	}
}
