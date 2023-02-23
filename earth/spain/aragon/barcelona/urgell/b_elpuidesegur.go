package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄尔贝塞居尔ElpuidesegurBarony struct {
	feud.BaseBarony
}

var BElpuidesegur厄尔贝塞居尔 feud.Barony = &厄尔贝塞居尔ElpuidesegurBarony{}

func init() {
	f := BElpuidesegur厄尔贝塞居尔.(*厄尔贝塞居尔ElpuidesegurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elpuidesegur",
		TitleName: "厄尔贝塞居尔",
		TitleCode: "b_elpuidesegur",
	}
}
