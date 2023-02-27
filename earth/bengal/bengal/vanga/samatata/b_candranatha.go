package samatata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀罗那他CandranathaBarony struct {
	feud.BaseBarony
}

var BCandranatha旃陀罗那他 feud.Barony = &旃陀罗那他CandranathaBarony{}

func init() {
    f := BCandranatha旃陀罗那他.(*旃陀罗那他CandranathaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "candranatha",
		TitleName: "旃陀罗那他",
		TitleCode: "b_candranatha",
	}
}
