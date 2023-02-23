package mecca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎姆穆JmumumBarony struct {
	feud.BaseBarony
}

var BJmumum扎姆穆 feud.Barony = &扎姆穆JmumumBarony{}

func init() {
	f := BJmumum扎姆穆.(*扎姆穆JmumumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jmumum",
		TitleName: "扎姆穆",
		TitleCode: "b_jmumum",
	}
}
