package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄斯特罗斯OsterasBarony struct {
	feud.BaseBarony
}

var BOsteras厄斯特罗斯 feud.Barony = &厄斯特罗斯OsterasBarony{}

func init() {
	f := BOsteras厄斯特罗斯.(*厄斯特罗斯OsterasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osteras",
		TitleName: "厄斯特罗斯",
		TitleCode: "b_osteras",
	}
}
