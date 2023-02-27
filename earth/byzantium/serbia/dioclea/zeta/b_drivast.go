package zeta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德里瓦斯特DrivastBarony struct {
	feud.BaseBarony
}

var BDrivast德里瓦斯特 feud.Barony = &德里瓦斯特DrivastBarony{}

func init() {
    f := BDrivast德里瓦斯特.(*德里瓦斯特DrivastBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drivast",
		TitleName: "德里瓦斯特",
		TitleCode: "b_drivast",
	}
}
