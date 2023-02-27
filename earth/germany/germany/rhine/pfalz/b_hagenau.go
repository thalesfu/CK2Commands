package pfalz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格诺HagenauBarony struct {
	feud.BaseBarony
}

var BHagenau阿格诺 feud.Barony = &阿格诺HagenauBarony{}

func init() {
    f := BHagenau阿格诺.(*阿格诺HagenauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hagenau",
		TitleName: "阿格诺",
		TitleCode: "b_hagenau",
	}
}
