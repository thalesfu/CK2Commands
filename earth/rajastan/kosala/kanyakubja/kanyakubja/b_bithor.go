package kanyakubja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗豆罗BithorBarony struct {
	feud.BaseBarony
}

var BBithor毗豆罗 feud.Barony = &毗豆罗BithorBarony{}

func init() {
    f := BBithor毗豆罗.(*毗豆罗BithorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bithor",
		TitleName: "毗豆罗",
		TitleCode: "b_bithor",
	}
}
