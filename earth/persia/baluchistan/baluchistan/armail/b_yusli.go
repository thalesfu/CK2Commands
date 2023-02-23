package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 育里YusliBarony struct {
	feud.BaseBarony
}

var BYusli育里 feud.Barony = &育里YusliBarony{}

func init() {
	f := BYusli育里.(*育里YusliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yusli",
		TitleName: "育里",
		TitleCode: "b_yusli",
	}
}
