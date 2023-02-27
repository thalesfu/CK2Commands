package bamiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔_格拉哈拉Shar_i_gholgholaBarony struct {
	feud.BaseBarony
}

var BShar_i_gholghola沙尔_格拉哈拉 feud.Barony = &沙尔_格拉哈拉Shar_i_gholgholaBarony{}

func init() {
    f := BShar_i_gholghola沙尔_格拉哈拉.(*沙尔_格拉哈拉Shar_i_gholgholaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shar_i_gholghola",
		TitleName: "沙尔_格拉哈拉",
		TitleCode: "b_shar_i_gholghola",
	}
}
