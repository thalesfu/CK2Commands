package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BreifneCounty interface {
    feud.County
    BArdagh阿达() 	feud.Barony
    BCavan卡文() 	feud.Barony
    BDromahair德罗马黑尔() 	feud.Barony
    BDrumcliffe杜木里菲() 	feud.Barony
    BKells凯尔斯() 	feud.Barony
    BKilmore基尔莫雷() 	feud.Barony
    BLeitrim利特里姆() 	feud.Barony
    BLongford朗福德() 	feud.Barony
}

type 布雷夫讷BreifneCounty struct {
	feud.BaseCounty
	阿达Ardagh 	feud.Barony
	卡文Cavan 	feud.Barony
	德罗马黑尔Dromahair 	feud.Barony
	杜木里菲Drumcliffe 	feud.Barony
	凯尔斯Kells 	feud.Barony
	基尔莫雷Kilmore 	feud.Barony
	利特里姆Leitrim 	feud.Barony
	朗福德Longford 	feud.Barony
}

func (c *布雷夫讷BreifneCounty) BArdagh阿达() feud.Barony {
	return c.阿达Ardagh
}
    
func (c *布雷夫讷BreifneCounty) BCavan卡文() feud.Barony {
	return c.卡文Cavan
}
    
func (c *布雷夫讷BreifneCounty) BDromahair德罗马黑尔() feud.Barony {
	return c.德罗马黑尔Dromahair
}
    
func (c *布雷夫讷BreifneCounty) BDrumcliffe杜木里菲() feud.Barony {
	return c.杜木里菲Drumcliffe
}
    
func (c *布雷夫讷BreifneCounty) BKells凯尔斯() feud.Barony {
	return c.凯尔斯Kells
}
    
func (c *布雷夫讷BreifneCounty) BKilmore基尔莫雷() feud.Barony {
	return c.基尔莫雷Kilmore
}
    
func (c *布雷夫讷BreifneCounty) BLeitrim利特里姆() feud.Barony {
	return c.利特里姆Leitrim
}
    
func (c *布雷夫讷BreifneCounty) BLongford朗福德() feud.Barony {
	return c.朗福德Longford
}
    
var CBreifne布雷夫讷 BreifneCounty = &布雷夫讷BreifneCounty{}

func init() {
	f := CBreifne布雷夫讷.(*布雷夫讷BreifneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "8",
		Title:     "breifne",
		TitleName: "布雷夫讷",
		TitleCode: "c_breifne",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿达Ardagh = BArdagh阿达
	f.阿达Ardagh.SetParent(f)
	
	f.卡文Cavan = BCavan卡文
	f.卡文Cavan.SetParent(f)
	
	f.德罗马黑尔Dromahair = BDromahair德罗马黑尔
	f.德罗马黑尔Dromahair.SetParent(f)
	
	f.杜木里菲Drumcliffe = BDrumcliffe杜木里菲
	f.杜木里菲Drumcliffe.SetParent(f)
	
	f.凯尔斯Kells = BKells凯尔斯
	f.凯尔斯Kells.SetParent(f)
	
	f.基尔莫雷Kilmore = BKilmore基尔莫雷
	f.基尔莫雷Kilmore.SetParent(f)
	
	f.利特里姆Leitrim = BLeitrim利特里姆
	f.利特里姆Leitrim.SetParent(f)
	
	f.朗福德Longford = BLongford朗福德
	f.朗福德Longford.SetParent(f)
	
}
