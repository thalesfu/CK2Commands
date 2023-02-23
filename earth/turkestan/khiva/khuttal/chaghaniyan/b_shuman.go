package chaghaniyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 数瞒ShumanBarony struct {
	feud.BaseBarony
}

var BShuman数瞒 feud.Barony = &数瞒ShumanBarony{}

func init() {
	f := BShuman数瞒.(*数瞒ShumanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shuman",
		TitleName: "数瞒",
		TitleCode: "b_shuman",
	}
}
