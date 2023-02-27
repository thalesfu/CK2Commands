package izborsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩乔雷PechoryBarony struct {
	feud.BaseBarony
}

var BPechory佩乔雷 feud.Barony = &佩乔雷PechoryBarony{}

func init() {
    f := BPechory佩乔雷.(*佩乔雷PechoryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pechory",
		TitleName: "佩乔雷",
		TitleCode: "b_pechory",
	}
}
