package sudovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SudoviaCounty interface {
    feud.County
    BAugustavas奥古斯塔瓦斯() 	feud.Barony
    BHrodna格罗德诺() 	feud.Barony
    BLiskiava利斯基亚瓦() 	feud.Barony
    BMerkine梅尔基内() 	feud.Barony
    BRudamina鲁达米纳() 	feud.Barony
    BSeiniai塞尼艾() 	feud.Barony
    BSuvalkai苏瓦尔凯() 	feud.Barony
}

type 苏多维亚SudoviaCounty struct {
	feud.BaseCounty
	奥古斯塔瓦斯Augustavas 	feud.Barony
	格罗德诺Hrodna 	feud.Barony
	利斯基亚瓦Liskiava 	feud.Barony
	梅尔基内Merkine 	feud.Barony
	鲁达米纳Rudamina 	feud.Barony
	塞尼艾Seiniai 	feud.Barony
	苏瓦尔凯Suvalkai 	feud.Barony
}

func (c *苏多维亚SudoviaCounty) BAugustavas奥古斯塔瓦斯() feud.Barony {
	return c.奥古斯塔瓦斯Augustavas
}
    
func (c *苏多维亚SudoviaCounty) BHrodna格罗德诺() feud.Barony {
	return c.格罗德诺Hrodna
}
    
func (c *苏多维亚SudoviaCounty) BLiskiava利斯基亚瓦() feud.Barony {
	return c.利斯基亚瓦Liskiava
}
    
func (c *苏多维亚SudoviaCounty) BMerkine梅尔基内() feud.Barony {
	return c.梅尔基内Merkine
}
    
func (c *苏多维亚SudoviaCounty) BRudamina鲁达米纳() feud.Barony {
	return c.鲁达米纳Rudamina
}
    
func (c *苏多维亚SudoviaCounty) BSeiniai塞尼艾() feud.Barony {
	return c.塞尼艾Seiniai
}
    
func (c *苏多维亚SudoviaCounty) BSuvalkai苏瓦尔凯() feud.Barony {
	return c.苏瓦尔凯Suvalkai
}
    
var CSudovia苏多维亚 SudoviaCounty = &苏多维亚SudoviaCounty{}

func init() {
	f := CSudovia苏多维亚.(*苏多维亚SudoviaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "423",
		Title:     "sudovia",
		TitleName: "苏多维亚",
		TitleCode: "c_sudovia",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥古斯塔瓦斯Augustavas = BAugustavas奥古斯塔瓦斯
	f.奥古斯塔瓦斯Augustavas.SetParent(f)
	
	f.格罗德诺Hrodna = BHrodna格罗德诺
	f.格罗德诺Hrodna.SetParent(f)
	
	f.利斯基亚瓦Liskiava = BLiskiava利斯基亚瓦
	f.利斯基亚瓦Liskiava.SetParent(f)
	
	f.梅尔基内Merkine = BMerkine梅尔基内
	f.梅尔基内Merkine.SetParent(f)
	
	f.鲁达米纳Rudamina = BRudamina鲁达米纳
	f.鲁达米纳Rudamina.SetParent(f)
	
	f.塞尼艾Seiniai = BSeiniai塞尼艾
	f.塞尼艾Seiniai.SetParent(f)
	
	f.苏瓦尔凯Suvalkai = BSuvalkai苏瓦尔凯
	f.苏瓦尔凯Suvalkai.SetParent(f)
	
}
