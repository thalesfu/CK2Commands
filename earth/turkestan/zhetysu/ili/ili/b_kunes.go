package ili

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巩乃斯KunesBarony struct {
	feud.BaseBarony
}

var BKunes巩乃斯 feud.Barony = &巩乃斯KunesBarony{}

func init() {
    f := BKunes巩乃斯.(*巩乃斯KunesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunes",
		TitleName: "巩乃斯",
		TitleCode: "b_kunes",
	}
}
