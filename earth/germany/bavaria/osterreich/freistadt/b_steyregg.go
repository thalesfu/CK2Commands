package freistadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施泰尔埃格SteyreggBarony struct {
	feud.BaseBarony
}

var BSteyregg施泰尔埃格 feud.Barony = &施泰尔埃格SteyreggBarony{}

func init() {
    f := BSteyregg施泰尔埃格.(*施泰尔埃格SteyreggBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "steyregg",
		TitleName: "施泰尔埃格",
		TitleCode: "b_steyregg",
	}
}
