package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤尔赫ElyourhBarony struct {
	feud.BaseBarony
}

var BElyourh尤尔赫 feud.Barony = &尤尔赫ElyourhBarony{}

func init() {
    f := BElyourh尤尔赫.(*尤尔赫ElyourhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elyourh",
		TitleName: "尤尔赫",
		TitleCode: "b_elyourh",
	}
}
