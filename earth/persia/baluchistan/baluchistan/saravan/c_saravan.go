package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaravanCounty interface {
    feud.County
    BJaleq贾勒格() 	feud.Barony
    BKhash哈什() 	feud.Barony
    BPahrah拉哈雷() 	feud.Barony
    BPishin皮欣() 	feud.Barony
    BRasak拉斯克() 	feud.Barony
    BSangan桑甘() 	feud.Barony
    BSaravan萨拉万() 	feud.Barony
    BSarbaz萨尔巴兹() 	feud.Barony
}

type 萨拉万SaravanCounty struct {
	feud.BaseCounty
	贾勒格Jaleq 	feud.Barony
	哈什Khash 	feud.Barony
	拉哈雷Pahrah 	feud.Barony
	皮欣Pishin 	feud.Barony
	拉斯克Rasak 	feud.Barony
	桑甘Sangan 	feud.Barony
	萨拉万Saravan 	feud.Barony
	萨尔巴兹Sarbaz 	feud.Barony
}

func (c *萨拉万SaravanCounty) BJaleq贾勒格() feud.Barony {
	return c.贾勒格Jaleq
}
    
func (c *萨拉万SaravanCounty) BKhash哈什() feud.Barony {
	return c.哈什Khash
}
    
func (c *萨拉万SaravanCounty) BPahrah拉哈雷() feud.Barony {
	return c.拉哈雷Pahrah
}
    
func (c *萨拉万SaravanCounty) BPishin皮欣() feud.Barony {
	return c.皮欣Pishin
}
    
func (c *萨拉万SaravanCounty) BRasak拉斯克() feud.Barony {
	return c.拉斯克Rasak
}
    
func (c *萨拉万SaravanCounty) BSangan桑甘() feud.Barony {
	return c.桑甘Sangan
}
    
func (c *萨拉万SaravanCounty) BSaravan萨拉万() feud.Barony {
	return c.萨拉万Saravan
}
    
func (c *萨拉万SaravanCounty) BSarbaz萨尔巴兹() feud.Barony {
	return c.萨尔巴兹Sarbaz
}
    
var CSaravan萨拉万 SaravanCounty = &萨拉万SaravanCounty{}

func init() {
	f := CSaravan萨拉万.(*萨拉万SaravanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "894",
		Title:     "saravan",
		TitleName: "萨拉万",
		TitleCode: "c_saravan",
		Baronies:  map[string]feud.Barony{},
	}

	f.贾勒格Jaleq = BJaleq贾勒格
	f.贾勒格Jaleq.SetParent(f)
	
	f.哈什Khash = BKhash哈什
	f.哈什Khash.SetParent(f)
	
	f.拉哈雷Pahrah = BPahrah拉哈雷
	f.拉哈雷Pahrah.SetParent(f)
	
	f.皮欣Pishin = BPishin皮欣
	f.皮欣Pishin.SetParent(f)
	
	f.拉斯克Rasak = BRasak拉斯克
	f.拉斯克Rasak.SetParent(f)
	
	f.桑甘Sangan = BSangan桑甘
	f.桑甘Sangan.SetParent(f)
	
	f.萨拉万Saravan = BSaravan萨拉万
	f.萨拉万Saravan.SetParent(f)
	
	f.萨尔巴兹Sarbaz = BSarbaz萨尔巴兹
	f.萨尔巴兹Sarbaz.SetParent(f)
	
}
