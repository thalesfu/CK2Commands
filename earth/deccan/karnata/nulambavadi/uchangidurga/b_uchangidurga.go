package uchangidurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 郁昌耆突伽UchangidurgaBarony struct {
	feud.BaseBarony
}

var BUchangidurga郁昌耆突伽 feud.Barony = &郁昌耆突伽UchangidurgaBarony{}

func init() {
	f := BUchangidurga郁昌耆突伽.(*郁昌耆突伽UchangidurgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uchangidurga",
		TitleName: "郁昌耆突伽",
		TitleCode: "b_uchangidurga",
	}
}
