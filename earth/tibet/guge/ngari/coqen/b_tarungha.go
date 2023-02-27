package coqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特荣海TarunghaBarony struct {
	feud.BaseBarony
}

var BTarungha特荣海 feud.Barony = &特荣海TarunghaBarony{}

func init() {
    f := BTarungha特荣海.(*特荣海TarunghaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarungha",
		TitleName: "特荣海",
		TitleCode: "b_tarungha",
	}
}
