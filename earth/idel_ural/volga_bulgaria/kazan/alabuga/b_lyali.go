package alabuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利亚利LyaliBarony struct {
	feud.BaseBarony
}

var BLyali利亚利 feud.Barony = &利亚利LyaliBarony{}

func init() {
    f := BLyali利亚利.(*利亚利LyaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyali",
		TitleName: "利亚利",
		TitleCode: "b_lyali",
	}
}
