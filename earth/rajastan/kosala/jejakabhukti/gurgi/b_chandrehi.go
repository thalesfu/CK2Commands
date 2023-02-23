package gurgi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀梨醯ChandrehiBarony struct {
	feud.BaseBarony
}

var BChandrehi旃陀梨醯 feud.Barony = &旃陀梨醯ChandrehiBarony{}

func init() {
	f := BChandrehi旃陀梨醯.(*旃陀梨醯ChandrehiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chandrehi",
		TitleName: "旃陀梨醯",
		TitleCode: "b_chandrehi",
	}
}
