package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱姆LymeBarony struct {
	feud.BaseBarony
}

var BLyme莱姆 feud.Barony = &莱姆LymeBarony{}

func init() {
    f := BLyme莱姆.(*莱姆LymeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyme",
		TitleName: "莱姆",
		TitleCode: "b_lyme",
	}
}
