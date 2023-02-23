package onega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OnegaCounty interface {
	feud.County
	BAunus奥努斯() feud.Barony
	BKondopoga孔多波加() feud.Barony
	BMedvezhyegorsk卡尔胡梅基() feud.Barony
	BPetrozavodsk彼得罗斯科伊() feud.Barony
	BPryazha普里亚扎() feud.Barony
	BPudoga普多加() feud.Barony
}

type 奥洛涅茨OnegaCounty struct {
	feud.BaseCounty
	奥努斯Aunus            feud.Barony
	孔多波加Kondopoga       feud.Barony
	卡尔胡梅基Medvezhyegorsk feud.Barony
	彼得罗斯科伊Petrozavodsk  feud.Barony
	普里亚扎Pryazha         feud.Barony
	普多加Pudoga           feud.Barony
}

func (c *奥洛涅茨OnegaCounty) BAunus奥努斯() feud.Barony {
	return c.奥努斯Aunus
}

func (c *奥洛涅茨OnegaCounty) BKondopoga孔多波加() feud.Barony {
	return c.孔多波加Kondopoga
}

func (c *奥洛涅茨OnegaCounty) BMedvezhyegorsk卡尔胡梅基() feud.Barony {
	return c.卡尔胡梅基Medvezhyegorsk
}

func (c *奥洛涅茨OnegaCounty) BPetrozavodsk彼得罗斯科伊() feud.Barony {
	return c.彼得罗斯科伊Petrozavodsk
}

func (c *奥洛涅茨OnegaCounty) BPryazha普里亚扎() feud.Barony {
	return c.普里亚扎Pryazha
}

func (c *奥洛涅茨OnegaCounty) BPudoga普多加() feud.Barony {
	return c.普多加Pudoga
}

var COnega奥洛涅茨 OnegaCounty = &奥洛涅茨OnegaCounty{}

func init() {
	f := COnega奥洛涅茨.(*奥洛涅茨OnegaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "393",
		Title:     "onega",
		TitleName: "奥洛涅茨",
		TitleCode: "c_onega",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥努斯Aunus = BAunus奥努斯
	f.奥努斯Aunus.SetParent(f)

	f.孔多波加Kondopoga = BKondopoga孔多波加
	f.孔多波加Kondopoga.SetParent(f)

	f.卡尔胡梅基Medvezhyegorsk = BMedvezhyegorsk卡尔胡梅基
	f.卡尔胡梅基Medvezhyegorsk.SetParent(f)

	f.彼得罗斯科伊Petrozavodsk = BPetrozavodsk彼得罗斯科伊
	f.彼得罗斯科伊Petrozavodsk.SetParent(f)

	f.普里亚扎Pryazha = BPryazha普里亚扎
	f.普里亚扎Pryazha.SetParent(f)

	f.普多加Pudoga = BPudoga普多加
	f.普多加Pudoga.SetParent(f)

}
