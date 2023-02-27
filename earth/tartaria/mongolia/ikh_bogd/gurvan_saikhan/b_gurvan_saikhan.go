package gurvan_saikhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔班赛罕Gurvan_saikhanBarony struct {
	feud.BaseBarony
}

var BGurvan_saikhan古尔班赛罕 feud.Barony = &古尔班赛罕Gurvan_saikhanBarony{}

func init() {
    f := BGurvan_saikhan古尔班赛罕.(*古尔班赛罕Gurvan_saikhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurvan_saikhan",
		TitleName: "古尔班赛罕",
		TitleCode: "b_gurvan_saikhan",
	}
}
