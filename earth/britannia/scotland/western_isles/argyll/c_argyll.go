package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArgyllCounty interface {
    feud.County
    BArdchattan艾彻坦() 	feud.Barony
    BArgh阿尔赫() 	feud.Barony
    BDunollie邓利() 	feud.Barony
    BDunstaffnage邓斯塔夫内奇() 	feud.Barony
    BInverary因弗拉里() 	feud.Barony
    BKilmun基尔曼() 	feud.Barony
    BLoch_awe奥湖() 	feud.Barony
    BSt_moluag圣莫洛格() 	feud.Barony
    BSween斯文() 	feud.Barony
}

type 阿盖尔ArgyllCounty struct {
	feud.BaseCounty
	艾彻坦Ardchattan 	feud.Barony
	阿尔赫Argh 	feud.Barony
	邓利Dunollie 	feud.Barony
	邓斯塔夫内奇Dunstaffnage 	feud.Barony
	因弗拉里Inverary 	feud.Barony
	基尔曼Kilmun 	feud.Barony
	奥湖Loch_awe 	feud.Barony
	圣莫洛格St_moluag 	feud.Barony
	斯文Sween 	feud.Barony
}

func (c *阿盖尔ArgyllCounty) BArdchattan艾彻坦() feud.Barony {
	return c.艾彻坦Ardchattan
}
    
func (c *阿盖尔ArgyllCounty) BArgh阿尔赫() feud.Barony {
	return c.阿尔赫Argh
}
    
func (c *阿盖尔ArgyllCounty) BDunollie邓利() feud.Barony {
	return c.邓利Dunollie
}
    
func (c *阿盖尔ArgyllCounty) BDunstaffnage邓斯塔夫内奇() feud.Barony {
	return c.邓斯塔夫内奇Dunstaffnage
}
    
func (c *阿盖尔ArgyllCounty) BInverary因弗拉里() feud.Barony {
	return c.因弗拉里Inverary
}
    
func (c *阿盖尔ArgyllCounty) BKilmun基尔曼() feud.Barony {
	return c.基尔曼Kilmun
}
    
func (c *阿盖尔ArgyllCounty) BLoch_awe奥湖() feud.Barony {
	return c.奥湖Loch_awe
}
    
func (c *阿盖尔ArgyllCounty) BSt_moluag圣莫洛格() feud.Barony {
	return c.圣莫洛格St_moluag
}
    
func (c *阿盖尔ArgyllCounty) BSween斯文() feud.Barony {
	return c.斯文Sween
}
    
var CArgyll阿盖尔 ArgyllCounty = &阿盖尔ArgyllCounty{}

func init() {
	f := CArgyll阿盖尔.(*阿盖尔ArgyllCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "45",
		Title:     "argyll",
		TitleName: "阿盖尔",
		TitleCode: "c_argyll",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾彻坦Ardchattan = BArdchattan艾彻坦
	f.艾彻坦Ardchattan.SetParent(f)
	
	f.阿尔赫Argh = BArgh阿尔赫
	f.阿尔赫Argh.SetParent(f)
	
	f.邓利Dunollie = BDunollie邓利
	f.邓利Dunollie.SetParent(f)
	
	f.邓斯塔夫内奇Dunstaffnage = BDunstaffnage邓斯塔夫内奇
	f.邓斯塔夫内奇Dunstaffnage.SetParent(f)
	
	f.因弗拉里Inverary = BInverary因弗拉里
	f.因弗拉里Inverary.SetParent(f)
	
	f.基尔曼Kilmun = BKilmun基尔曼
	f.基尔曼Kilmun.SetParent(f)
	
	f.奥湖Loch_awe = BLoch_awe奥湖
	f.奥湖Loch_awe.SetParent(f)
	
	f.圣莫洛格St_moluag = BSt_moluag圣莫洛格
	f.圣莫洛格St_moluag.SetParent(f)
	
	f.斯文Sween = BSween斯文
	f.斯文Sween.SetParent(f)
	
}
