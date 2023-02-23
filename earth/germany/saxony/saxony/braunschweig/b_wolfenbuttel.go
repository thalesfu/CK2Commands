package braunschweig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔芬比特尔WolfenbuttelBarony struct {
	feud.BaseBarony
}

var BWolfenbuttel沃尔芬比特尔 feud.Barony = &沃尔芬比特尔WolfenbuttelBarony{}

func init() {
	f := BWolfenbuttel沃尔芬比特尔.(*沃尔芬比特尔WolfenbuttelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wolfenbuttel",
		TitleName: "沃尔芬比特尔",
		TitleCode: "b_wolfenbuttel",
	}
}
