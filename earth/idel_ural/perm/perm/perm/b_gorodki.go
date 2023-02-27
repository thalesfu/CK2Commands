package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格罗德基GorodkiBarony struct {
	feud.BaseBarony
}

var BGorodki格罗德基 feud.Barony = &格罗德基GorodkiBarony{}

func init() {
    f := BGorodki格罗德基.(*格罗德基GorodkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorodki",
		TitleName: "格罗德基",
		TitleCode: "b_gorodki",
	}
}
