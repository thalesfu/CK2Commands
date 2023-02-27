package yarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸦儿看YarkandBarony struct {
	feud.BaseBarony
}

var BYarkand鸦儿看 feud.Barony = &鸦儿看YarkandBarony{}

func init() {
    f := BYarkand鸦儿看.(*鸦儿看YarkandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yarkand",
		TitleName: "鸦儿看",
		TitleCode: "b_yarkand",
	}
}
