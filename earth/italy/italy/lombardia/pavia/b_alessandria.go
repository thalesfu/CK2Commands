package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚历山德里亚AlessandriaBarony struct {
	feud.BaseBarony
}

var BAlessandria亚历山德里亚 feud.Barony = &亚历山德里亚AlessandriaBarony{}

func init() {
	f := BAlessandria亚历山德里亚.(*亚历山德里亚AlessandriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alessandria",
		TitleName: "亚历山德里亚",
		TitleCode: "b_alessandria",
	}
}
