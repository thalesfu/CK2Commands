package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼尔姆PermBarony struct {
	feud.BaseBarony
}

var BPerm彼尔姆 feud.Barony = &彼尔姆PermBarony{}

func init() {
    f := BPerm彼尔姆.(*彼尔姆PermBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perm",
		TitleName: "彼尔姆",
		TitleCode: "b_perm",
	}
}
