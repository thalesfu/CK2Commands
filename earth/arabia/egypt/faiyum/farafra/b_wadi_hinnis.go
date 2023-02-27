package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欣尼斯谷Wadi_hinnisBarony struct {
	feud.BaseBarony
}

var BWadi_hinnis欣尼斯谷 feud.Barony = &欣尼斯谷Wadi_hinnisBarony{}

func init() {
    f := BWadi_hinnis欣尼斯谷.(*欣尼斯谷Wadi_hinnisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wadi_hinnis",
		TitleName: "欣尼斯谷",
		TitleCode: "b_wadi_hinnis",
	}
}
