package konjikala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KonjikalaCounty interface {
    feud.County
    BAbiward阿比瓦尔德() 	feud.Barony
    BGokdepe吉奥克杰佩() 	feud.Barony
    BKonjikala孔吉卡拉() 	feud.Barony
    BKyzylarvat克孜勒阿尔瓦特() 	feud.Barony
    BNisa尼萨() 	feud.Barony
    BSarahs谢拉赫斯() 	feud.Barony
    BUlugdepe乌卢格_杰佩() 	feud.Barony
}

type 孔吉卡拉KonjikalaCounty struct {
	feud.BaseCounty
	阿比瓦尔德Abiward 	feud.Barony
	吉奥克杰佩Gokdepe 	feud.Barony
	孔吉卡拉Konjikala 	feud.Barony
	克孜勒阿尔瓦特Kyzylarvat 	feud.Barony
	尼萨Nisa 	feud.Barony
	谢拉赫斯Sarahs 	feud.Barony
	乌卢格_杰佩Ulugdepe 	feud.Barony
}

func (c *孔吉卡拉KonjikalaCounty) BAbiward阿比瓦尔德() feud.Barony {
	return c.阿比瓦尔德Abiward
}
    
func (c *孔吉卡拉KonjikalaCounty) BGokdepe吉奥克杰佩() feud.Barony {
	return c.吉奥克杰佩Gokdepe
}
    
func (c *孔吉卡拉KonjikalaCounty) BKonjikala孔吉卡拉() feud.Barony {
	return c.孔吉卡拉Konjikala
}
    
func (c *孔吉卡拉KonjikalaCounty) BKyzylarvat克孜勒阿尔瓦特() feud.Barony {
	return c.克孜勒阿尔瓦特Kyzylarvat
}
    
func (c *孔吉卡拉KonjikalaCounty) BNisa尼萨() feud.Barony {
	return c.尼萨Nisa
}
    
func (c *孔吉卡拉KonjikalaCounty) BSarahs谢拉赫斯() feud.Barony {
	return c.谢拉赫斯Sarahs
}
    
func (c *孔吉卡拉KonjikalaCounty) BUlugdepe乌卢格_杰佩() feud.Barony {
	return c.乌卢格_杰佩Ulugdepe
}
    
var CKonjikala孔吉卡拉 KonjikalaCounty = &孔吉卡拉KonjikalaCounty{}

func init() {
	f := CKonjikala孔吉卡拉.(*孔吉卡拉KonjikalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "629",
		Title:     "konjikala",
		TitleName: "孔吉卡拉",
		TitleCode: "c_konjikala",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿比瓦尔德Abiward = BAbiward阿比瓦尔德
	f.阿比瓦尔德Abiward.SetParent(f)
	
	f.吉奥克杰佩Gokdepe = BGokdepe吉奥克杰佩
	f.吉奥克杰佩Gokdepe.SetParent(f)
	
	f.孔吉卡拉Konjikala = BKonjikala孔吉卡拉
	f.孔吉卡拉Konjikala.SetParent(f)
	
	f.克孜勒阿尔瓦特Kyzylarvat = BKyzylarvat克孜勒阿尔瓦特
	f.克孜勒阿尔瓦特Kyzylarvat.SetParent(f)
	
	f.尼萨Nisa = BNisa尼萨
	f.尼萨Nisa.SetParent(f)
	
	f.谢拉赫斯Sarahs = BSarahs谢拉赫斯
	f.谢拉赫斯Sarahs.SetParent(f)
	
	f.乌卢格_杰佩Ulugdepe = BUlugdepe乌卢格_杰佩
	f.乌卢格_杰佩Ulugdepe.SetParent(f)
	
}
