package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉凯戴孟尼亚LacedaemoniaBarony struct {
	feud.BaseBarony
}

var BLacedaemonia拉凯戴孟尼亚 feud.Barony = &拉凯戴孟尼亚LacedaemoniaBarony{}

func init() {
    f := BLacedaemonia拉凯戴孟尼亚.(*拉凯戴孟尼亚LacedaemoniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lacedaemonia",
		TitleName: "拉凯戴孟尼亚",
		TitleCode: "b_lacedaemonia",
	}
}
