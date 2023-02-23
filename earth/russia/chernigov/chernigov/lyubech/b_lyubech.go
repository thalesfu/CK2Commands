package lyubech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳别奇LyubechBarony struct {
	feud.BaseBarony
}

var BLyubech柳别奇 feud.Barony = &柳别奇LyubechBarony{}

func init() {
	f := BLyubech柳别奇.(*柳别奇LyubechBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyubech",
		TitleName: "柳别奇",
		TitleCode: "b_lyubech",
	}
}
