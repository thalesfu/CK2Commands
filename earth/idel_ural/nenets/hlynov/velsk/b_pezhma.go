package velsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩日马PezhmaBarony struct {
	feud.BaseBarony
}

var BPezhma佩日马 feud.Barony = &佩日马PezhmaBarony{}

func init() {
    f := BPezhma佩日马.(*佩日马PezhmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pezhma",
		TitleName: "佩日马",
		TitleCode: "b_pezhma",
	}
}
