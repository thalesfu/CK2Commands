package lhatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LhatseCounty interface {
	feud.County
	BDeleg德来() feud.Barony
	BDrampagyang仲巴江() feud.Barony
	BGyangbumoche江蚌摩齐() feud.Barony
	BKagar卡嘎() feud.Barony
	BLhatse拉孜() feud.Barony
	BPuncogling彭措林() feud.Barony
	BQuxar曲下() feud.Barony
}

type 拉孜LhatseCounty struct {
	feud.BaseCounty
	德来Deleg          feud.Barony
	仲巴江Drampagyang   feud.Barony
	江蚌摩齐Gyangbumoche feud.Barony
	卡嘎Kagar          feud.Barony
	拉孜Lhatse         feud.Barony
	彭措林Puncogling    feud.Barony
	曲下Quxar          feud.Barony
}

func (c *拉孜LhatseCounty) BDeleg德来() feud.Barony {
	return c.德来Deleg
}

func (c *拉孜LhatseCounty) BDrampagyang仲巴江() feud.Barony {
	return c.仲巴江Drampagyang
}

func (c *拉孜LhatseCounty) BGyangbumoche江蚌摩齐() feud.Barony {
	return c.江蚌摩齐Gyangbumoche
}

func (c *拉孜LhatseCounty) BKagar卡嘎() feud.Barony {
	return c.卡嘎Kagar
}

func (c *拉孜LhatseCounty) BLhatse拉孜() feud.Barony {
	return c.拉孜Lhatse
}

func (c *拉孜LhatseCounty) BPuncogling彭措林() feud.Barony {
	return c.彭措林Puncogling
}

func (c *拉孜LhatseCounty) BQuxar曲下() feud.Barony {
	return c.曲下Quxar
}

var CLhatse拉孜 LhatseCounty = &拉孜LhatseCounty{}

func init() {
	f := CLhatse拉孜.(*拉孜LhatseCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1561",
		Title:     "lhatse",
		TitleName: "拉孜",
		TitleCode: "c_lhatse",
		Baronies:  map[string]feud.Barony{},
	}

	f.德来Deleg = BDeleg德来
	f.德来Deleg.SetParent(f)

	f.仲巴江Drampagyang = BDrampagyang仲巴江
	f.仲巴江Drampagyang.SetParent(f)

	f.江蚌摩齐Gyangbumoche = BGyangbumoche江蚌摩齐
	f.江蚌摩齐Gyangbumoche.SetParent(f)

	f.卡嘎Kagar = BKagar卡嘎
	f.卡嘎Kagar.SetParent(f)

	f.拉孜Lhatse = BLhatse拉孜
	f.拉孜Lhatse.SetParent(f)

	f.彭措林Puncogling = BPuncogling彭措林
	f.彭措林Puncogling.SetParent(f)

	f.曲下Quxar = BQuxar曲下
	f.曲下Quxar.SetParent(f)

}
