package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OromiehCounty interface {
    feud.County
    BAvajiq阿瓦吉克() 	feud.Barony
    BChaldiran恰尔德兰() 	feud.Barony
    BCharekelisa察雷克里撒() 	feud.Barony
    BKhoy霍伊() 	feud.Barony
    BMahabad马哈巴德() 	feud.Barony
    BOroumieh乌鲁米耶() 	feud.Barony
    BSalmas萨勒马斯() 	feud.Barony
    BTakhtesoleyman塔赫特苏莱曼() 	feud.Barony
}

type 埃尔比勒OromiehCounty struct {
	feud.BaseCounty
	阿瓦吉克Avajiq 	feud.Barony
	恰尔德兰Chaldiran 	feud.Barony
	察雷克里撒Charekelisa 	feud.Barony
	霍伊Khoy 	feud.Barony
	马哈巴德Mahabad 	feud.Barony
	乌鲁米耶Oroumieh 	feud.Barony
	萨勒马斯Salmas 	feud.Barony
	塔赫特苏莱曼Takhtesoleyman 	feud.Barony
}

func (c *埃尔比勒OromiehCounty) BAvajiq阿瓦吉克() feud.Barony {
	return c.阿瓦吉克Avajiq
}
    
func (c *埃尔比勒OromiehCounty) BChaldiran恰尔德兰() feud.Barony {
	return c.恰尔德兰Chaldiran
}
    
func (c *埃尔比勒OromiehCounty) BCharekelisa察雷克里撒() feud.Barony {
	return c.察雷克里撒Charekelisa
}
    
func (c *埃尔比勒OromiehCounty) BKhoy霍伊() feud.Barony {
	return c.霍伊Khoy
}
    
func (c *埃尔比勒OromiehCounty) BMahabad马哈巴德() feud.Barony {
	return c.马哈巴德Mahabad
}
    
func (c *埃尔比勒OromiehCounty) BOroumieh乌鲁米耶() feud.Barony {
	return c.乌鲁米耶Oroumieh
}
    
func (c *埃尔比勒OromiehCounty) BSalmas萨勒马斯() feud.Barony {
	return c.萨勒马斯Salmas
}
    
func (c *埃尔比勒OromiehCounty) BTakhtesoleyman塔赫特苏莱曼() feud.Barony {
	return c.塔赫特苏莱曼Takhtesoleyman
}
    
var COromieh埃尔比勒 OromiehCounty = &埃尔比勒OromiehCounty{}

func init() {
	f := COromieh埃尔比勒.(*埃尔比勒OromiehCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "685",
		Title:     "oromieh",
		TitleName: "埃尔比勒",
		TitleCode: "c_oromieh",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿瓦吉克Avajiq = BAvajiq阿瓦吉克
	f.阿瓦吉克Avajiq.SetParent(f)
	
	f.恰尔德兰Chaldiran = BChaldiran恰尔德兰
	f.恰尔德兰Chaldiran.SetParent(f)
	
	f.察雷克里撒Charekelisa = BCharekelisa察雷克里撒
	f.察雷克里撒Charekelisa.SetParent(f)
	
	f.霍伊Khoy = BKhoy霍伊
	f.霍伊Khoy.SetParent(f)
	
	f.马哈巴德Mahabad = BMahabad马哈巴德
	f.马哈巴德Mahabad.SetParent(f)
	
	f.乌鲁米耶Oroumieh = BOroumieh乌鲁米耶
	f.乌鲁米耶Oroumieh.SetParent(f)
	
	f.萨勒马斯Salmas = BSalmas萨勒马斯
	f.萨勒马斯Salmas.SetParent(f)
	
	f.塔赫特苏莱曼Takhtesoleyman = BTakhtesoleyman塔赫特苏莱曼
	f.塔赫特苏莱曼Takhtesoleyman.SetParent(f)
	
}
