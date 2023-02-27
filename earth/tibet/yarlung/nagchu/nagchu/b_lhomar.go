package nagchu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗玛LhomarBarony struct {
	feud.BaseBarony
}

var BLhomar罗玛 feud.Barony = &罗玛LhomarBarony{}

func init() {
    f := BLhomar罗玛.(*罗玛LhomarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lhomar",
		TitleName: "罗玛",
		TitleCode: "b_lhomar",
	}
}
