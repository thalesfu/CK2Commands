package nurnberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽伦堡NurnbergBarony struct {
	feud.BaseBarony
}

var BNurnberg纽伦堡 feud.Barony = &纽伦堡NurnbergBarony{}

func init() {
    f := BNurnberg纽伦堡.(*纽伦堡NurnbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nurnberg",
		TitleName: "纽伦堡",
		TitleCode: "b_nurnberg",
	}
}
