package khuttal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 久越得犍QobadhiyanBarony struct {
	feud.BaseBarony
}

var BQobadhiyan久越得犍 feud.Barony = &久越得犍QobadhiyanBarony{}

func init() {
	f := BQobadhiyan久越得犍.(*久越得犍QobadhiyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qobadhiyan",
		TitleName: "久越得犍",
		TitleCode: "b_qobadhiyan",
	}
}
