package lyncestis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩拉戈尼亚PelagoniaBarony struct {
	feud.BaseBarony
}

var BPelagonia佩拉戈尼亚 feud.Barony = &佩拉戈尼亚PelagoniaBarony{}

func init() {
    f := BPelagonia佩拉戈尼亚.(*佩拉戈尼亚PelagoniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pelagonia",
		TitleName: "佩拉戈尼亚",
		TitleCode: "b_pelagonia",
	}
}
