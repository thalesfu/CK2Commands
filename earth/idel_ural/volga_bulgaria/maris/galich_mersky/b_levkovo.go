package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列夫科沃LevkovoBarony struct {
	feud.BaseBarony
}

var BLevkovo列夫科沃 feud.Barony = &列夫科沃LevkovoBarony{}

func init() {
    f := BLevkovo列夫科沃.(*列夫科沃LevkovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "levkovo",
		TitleName: "列夫科沃",
		TitleCode: "b_levkovo",
	}
}
