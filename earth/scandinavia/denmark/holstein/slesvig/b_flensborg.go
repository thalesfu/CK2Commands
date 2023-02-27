package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗伦斯堡FlensborgBarony struct {
	feud.BaseBarony
}

var BFlensborg弗伦斯堡 feud.Barony = &弗伦斯堡FlensborgBarony{}

func init() {
    f := BFlensborg弗伦斯堡.(*弗伦斯堡FlensborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "flensborg",
		TitleName: "弗伦斯堡",
		TitleCode: "b_flensborg",
	}
}
