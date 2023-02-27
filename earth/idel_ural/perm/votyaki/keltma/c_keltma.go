package keltma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KeltmaCounty interface {
    feud.County
    BKadamskoye卡达姆斯科耶() 	feud.Barony
    BKeltma克利特马() 	feud.Barony
    BKerchomya克尔乔姆亚() 	feud.Barony
    BNem涅姆() 	feud.Barony
    BOkos奥科斯() 	feud.Barony
    BSmolyanka斯莫良卡() 	feud.Barony
    BYugyd_yag尤格德亚格() 	feud.Barony
}

type 克利特马KeltmaCounty struct {
	feud.BaseCounty
	卡达姆斯科耶Kadamskoye 	feud.Barony
	克利特马Keltma 	feud.Barony
	克尔乔姆亚Kerchomya 	feud.Barony
	涅姆Nem 	feud.Barony
	奥科斯Okos 	feud.Barony
	斯莫良卡Smolyanka 	feud.Barony
	尤格德亚格Yugyd_yag 	feud.Barony
}

func (c *克利特马KeltmaCounty) BKadamskoye卡达姆斯科耶() feud.Barony {
	return c.卡达姆斯科耶Kadamskoye
}
    
func (c *克利特马KeltmaCounty) BKeltma克利特马() feud.Barony {
	return c.克利特马Keltma
}
    
func (c *克利特马KeltmaCounty) BKerchomya克尔乔姆亚() feud.Barony {
	return c.克尔乔姆亚Kerchomya
}
    
func (c *克利特马KeltmaCounty) BNem涅姆() feud.Barony {
	return c.涅姆Nem
}
    
func (c *克利特马KeltmaCounty) BOkos奥科斯() feud.Barony {
	return c.奥科斯Okos
}
    
func (c *克利特马KeltmaCounty) BSmolyanka斯莫良卡() feud.Barony {
	return c.斯莫良卡Smolyanka
}
    
func (c *克利特马KeltmaCounty) BYugyd_yag尤格德亚格() feud.Barony {
	return c.尤格德亚格Yugyd_yag
}
    
var CKeltma克利特马 KeltmaCounty = &克利特马KeltmaCounty{}

func init() {
	f := CKeltma克利特马.(*克利特马KeltmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1842",
		Title:     "keltma",
		TitleName: "克利特马",
		TitleCode: "c_keltma",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡达姆斯科耶Kadamskoye = BKadamskoye卡达姆斯科耶
	f.卡达姆斯科耶Kadamskoye.SetParent(f)
	
	f.克利特马Keltma = BKeltma克利特马
	f.克利特马Keltma.SetParent(f)
	
	f.克尔乔姆亚Kerchomya = BKerchomya克尔乔姆亚
	f.克尔乔姆亚Kerchomya.SetParent(f)
	
	f.涅姆Nem = BNem涅姆
	f.涅姆Nem.SetParent(f)
	
	f.奥科斯Okos = BOkos奥科斯
	f.奥科斯Okos.SetParent(f)
	
	f.斯莫良卡Smolyanka = BSmolyanka斯莫良卡
	f.斯莫良卡Smolyanka.SetParent(f)
	
	f.尤格德亚格Yugyd_yag = BYugyd_yag尤格德亚格
	f.尤格德亚格Yugyd_yag.SetParent(f)
	
}
