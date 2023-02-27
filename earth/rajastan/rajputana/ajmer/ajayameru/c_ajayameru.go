package ajayameru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AjayameruCounty interface {
    feud.County
    BAjayameru阿阇耶迷卢() 	feud.Barony
    BDhanop檀努() 	feud.Barony
    BLawah罗瓦哈() 	feud.Barony
    BNareli那丽利() 	feud.Barony
    BPushkar布色羯逻() 	feud.Barony
    BShakambhari舍剑婆梨() 	feud.Barony
    BTaragarh多罗姞利呬() 	feud.Barony
}

type 阿阇耶迷卢AjayameruCounty struct {
	feud.BaseCounty
	阿阇耶迷卢Ajayameru 	feud.Barony
	檀努Dhanop 	feud.Barony
	罗瓦哈Lawah 	feud.Barony
	那丽利Nareli 	feud.Barony
	布色羯逻Pushkar 	feud.Barony
	舍剑婆梨Shakambhari 	feud.Barony
	多罗姞利呬Taragarh 	feud.Barony
}

func (c *阿阇耶迷卢AjayameruCounty) BAjayameru阿阇耶迷卢() feud.Barony {
	return c.阿阇耶迷卢Ajayameru
}
    
func (c *阿阇耶迷卢AjayameruCounty) BDhanop檀努() feud.Barony {
	return c.檀努Dhanop
}
    
func (c *阿阇耶迷卢AjayameruCounty) BLawah罗瓦哈() feud.Barony {
	return c.罗瓦哈Lawah
}
    
func (c *阿阇耶迷卢AjayameruCounty) BNareli那丽利() feud.Barony {
	return c.那丽利Nareli
}
    
func (c *阿阇耶迷卢AjayameruCounty) BPushkar布色羯逻() feud.Barony {
	return c.布色羯逻Pushkar
}
    
func (c *阿阇耶迷卢AjayameruCounty) BShakambhari舍剑婆梨() feud.Barony {
	return c.舍剑婆梨Shakambhari
}
    
func (c *阿阇耶迷卢AjayameruCounty) BTaragarh多罗姞利呬() feud.Barony {
	return c.多罗姞利呬Taragarh
}
    
var CAjayameru阿阇耶迷卢 AjayameruCounty = &阿阇耶迷卢AjayameruCounty{}

func init() {
	f := CAjayameru阿阇耶迷卢.(*阿阇耶迷卢AjayameruCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1346",
		Title:     "ajayameru",
		TitleName: "阿阇耶迷卢",
		TitleCode: "c_ajayameru",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿阇耶迷卢Ajayameru = BAjayameru阿阇耶迷卢
	f.阿阇耶迷卢Ajayameru.SetParent(f)
	
	f.檀努Dhanop = BDhanop檀努
	f.檀努Dhanop.SetParent(f)
	
	f.罗瓦哈Lawah = BLawah罗瓦哈
	f.罗瓦哈Lawah.SetParent(f)
	
	f.那丽利Nareli = BNareli那丽利
	f.那丽利Nareli.SetParent(f)
	
	f.布色羯逻Pushkar = BPushkar布色羯逻
	f.布色羯逻Pushkar.SetParent(f)
	
	f.舍剑婆梨Shakambhari = BShakambhari舍剑婆梨
	f.舍剑婆梨Shakambhari.SetParent(f)
	
	f.多罗姞利呬Taragarh = BTaragarh多罗姞利呬
	f.多罗姞利呬Taragarh.SetParent(f)
	
}
