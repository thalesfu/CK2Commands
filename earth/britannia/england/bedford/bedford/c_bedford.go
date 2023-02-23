package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BedfordCounty interface {
	feud.County
	BAshridge阿什里奇() feud.Barony
	BBedford贝德福德() feud.Barony
	BBerkhamsted伯克姆斯特德() feud.Barony
	BDunstable邓斯特布尔() feud.Barony
	BHertford赫特福德() feud.Barony
	BLuton卢顿() feud.Barony
	BWatford沃特福德() feud.Barony
}

type 贝德福德BedfordCounty struct {
	feud.BaseCounty
	阿什里奇Ashridge      feud.Barony
	贝德福德Bedford       feud.Barony
	伯克姆斯特德Berkhamsted feud.Barony
	邓斯特布尔Dunstable    feud.Barony
	赫特福德Hertford      feud.Barony
	卢顿Luton           feud.Barony
	沃特福德Watford       feud.Barony
}

func (c *贝德福德BedfordCounty) BAshridge阿什里奇() feud.Barony {
	return c.阿什里奇Ashridge
}

func (c *贝德福德BedfordCounty) BBedford贝德福德() feud.Barony {
	return c.贝德福德Bedford
}

func (c *贝德福德BedfordCounty) BBerkhamsted伯克姆斯特德() feud.Barony {
	return c.伯克姆斯特德Berkhamsted
}

func (c *贝德福德BedfordCounty) BDunstable邓斯特布尔() feud.Barony {
	return c.邓斯特布尔Dunstable
}

func (c *贝德福德BedfordCounty) BHertford赫特福德() feud.Barony {
	return c.赫特福德Hertford
}

func (c *贝德福德BedfordCounty) BLuton卢顿() feud.Barony {
	return c.卢顿Luton
}

func (c *贝德福德BedfordCounty) BWatford沃特福德() feud.Barony {
	return c.沃特福德Watford
}

var CBedford贝德福德 BedfordCounty = &贝德福德BedfordCounty{}

func init() {
	f := CBedford贝德福德.(*贝德福德BedfordCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "69",
		Title:     "bedford",
		TitleName: "贝德福德",
		TitleCode: "c_bedford",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿什里奇Ashridge = BAshridge阿什里奇
	f.阿什里奇Ashridge.SetParent(f)

	f.贝德福德Bedford = BBedford贝德福德
	f.贝德福德Bedford.SetParent(f)

	f.伯克姆斯特德Berkhamsted = BBerkhamsted伯克姆斯特德
	f.伯克姆斯特德Berkhamsted.SetParent(f)

	f.邓斯特布尔Dunstable = BDunstable邓斯特布尔
	f.邓斯特布尔Dunstable.SetParent(f)

	f.赫特福德Hertford = BHertford赫特福德
	f.赫特福德Hertford.SetParent(f)

	f.卢顿Luton = BLuton卢顿
	f.卢顿Luton.SetParent(f)

	f.沃特福德Watford = BWatford沃特福德
	f.沃特福德Watford.SetParent(f)

}
