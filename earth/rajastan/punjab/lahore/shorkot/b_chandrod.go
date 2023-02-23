package shorkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀卢陀ChandrodBarony struct {
	feud.BaseBarony
}

var BChandrod旃陀卢陀 feud.Barony = &旃陀卢陀ChandrodBarony{}

func init() {
	f := BChandrod旃陀卢陀.(*旃陀卢陀ChandrodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chandrod",
		TitleName: "旃陀卢陀",
		TitleCode: "b_chandrod",
	}
}
