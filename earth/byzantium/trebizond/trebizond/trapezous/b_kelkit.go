package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔基特KelkitBarony struct {
	feud.BaseBarony
}

var BKelkit凯尔基特 feud.Barony = &凯尔基特KelkitBarony{}

func init() {
	f := BKelkit凯尔基特.(*凯尔基特KelkitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kelkit",
		TitleName: "凯尔基特",
		TitleCode: "b_kelkit",
	}
}
