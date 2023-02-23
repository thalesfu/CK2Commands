package borovichi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳贝季诺LyubytinoBarony struct {
	feud.BaseBarony
}

var BLyubytino柳贝季诺 feud.Barony = &柳贝季诺LyubytinoBarony{}

func init() {
	f := BLyubytino柳贝季诺.(*柳贝季诺LyubytinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyubytino",
		TitleName: "柳贝季诺",
		TitleCode: "b_lyubytino",
	}
}
