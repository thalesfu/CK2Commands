package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于克明德UeckermundeBarony struct {
	feud.BaseBarony
}

var BUeckermunde于克明德 feud.Barony = &于克明德UeckermundeBarony{}

func init() {
    f := BUeckermunde于克明德.(*于克明德UeckermundeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ueckermunde",
		TitleName: "于克明德",
		TitleCode: "b_ueckermunde",
	}
}
