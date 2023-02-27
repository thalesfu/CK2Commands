package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪贝勒阿贝斯SidibelabbesBarony struct {
	feud.BaseBarony
}

var BSidibelabbes西迪贝勒阿贝斯 feud.Barony = &西迪贝勒阿贝斯SidibelabbesBarony{}

func init() {
    f := BSidibelabbes西迪贝勒阿贝斯.(*西迪贝勒阿贝斯SidibelabbesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidibelabbes",
		TitleName: "西迪贝勒阿贝斯",
		TitleCode: "b_sidibelabbes",
	}
}
