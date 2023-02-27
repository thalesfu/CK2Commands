package asirgarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇迦补罗JakpurBarony struct {
	feud.BaseBarony
}

var BJakpur阇迦补罗 feud.Barony = &阇迦补罗JakpurBarony{}

func init() {
    f := BJakpur阇迦补罗.(*阇迦补罗JakpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jakpur",
		TitleName: "阇迦补罗",
		TitleCode: "b_jakpur",
	}
}
