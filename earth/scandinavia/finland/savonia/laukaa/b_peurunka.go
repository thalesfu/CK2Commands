package laukaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩乌伦卡PeurunkaBarony struct {
	feud.BaseBarony
}

var BPeurunka佩乌伦卡 feud.Barony = &佩乌伦卡PeurunkaBarony{}

func init() {
    f := BPeurunka佩乌伦卡.(*佩乌伦卡PeurunkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peurunka",
		TitleName: "佩乌伦卡",
		TitleCode: "b_peurunka",
	}
}
