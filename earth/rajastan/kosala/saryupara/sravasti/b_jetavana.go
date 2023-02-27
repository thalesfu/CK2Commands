package sravasti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祇园JetavanaBarony struct {
	feud.BaseBarony
}

var BJetavana祇园 feud.Barony = &祇园JetavanaBarony{}

func init() {
    f := BJetavana祇园.(*祇园JetavanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jetavana",
		TitleName: "祇园",
		TitleCode: "b_jetavana",
	}
}
