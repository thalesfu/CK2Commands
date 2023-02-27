package kremenchuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克列缅丘格VolgakremenchukBarony struct {
	feud.BaseBarony
}

var BVolgakremenchuk克列缅丘格 feud.Barony = &克列缅丘格VolgakremenchukBarony{}

func init() {
    f := BVolgakremenchuk克列缅丘格.(*克列缅丘格VolgakremenchukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volgakremenchuk",
		TitleName: "克列缅丘格",
		TitleCode: "b_volgakremenchuk",
	}
}
