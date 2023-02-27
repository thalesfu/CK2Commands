package karnten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉根福特KlagenfurtBarony struct {
	feud.BaseBarony
}

var BKlagenfurt克拉根福特 feud.Barony = &克拉根福特KlagenfurtBarony{}

func init() {
    f := BKlagenfurt克拉根福特.(*克拉根福特KlagenfurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klagenfurt",
		TitleName: "克拉根福特",
		TitleCode: "b_klagenfurt",
	}
}
