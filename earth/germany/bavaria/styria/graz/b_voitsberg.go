package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福伊茨贝格VoitsbergBarony struct {
	feud.BaseBarony
}

var BVoitsberg福伊茨贝格 feud.Barony = &福伊茨贝格VoitsbergBarony{}

func init() {
	f := BVoitsberg福伊茨贝格.(*福伊茨贝格VoitsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voitsberg",
		TitleName: "福伊茨贝格",
		TitleCode: "b_voitsberg",
	}
}
