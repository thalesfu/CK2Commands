package yopurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YopurgaCounty interface {
    feud.County
    BBeijiang北降() 	feud.Barony
    BBiancheng遍城() 	feud.Barony
    BFanjiaxiaozhang范家小掌() 	feud.Barony
    BHoushan后善() 	feud.Barony
    BJiashi伽师() 	feud.Barony
    BPalichuang八里庄() 	feud.Barony
    BYopurga岳普湖() 	feud.Barony
}

type 岳普湖YopurgaCounty struct {
	feud.BaseCounty
	北降Beijiang 	feud.Barony
	遍城Biancheng 	feud.Barony
	范家小掌Fanjiaxiaozhang 	feud.Barony
	后善Houshan 	feud.Barony
	伽师Jiashi 	feud.Barony
	八里庄Palichuang 	feud.Barony
	岳普湖Yopurga 	feud.Barony
}

func (c *岳普湖YopurgaCounty) BBeijiang北降() feud.Barony {
	return c.北降Beijiang
}
    
func (c *岳普湖YopurgaCounty) BBiancheng遍城() feud.Barony {
	return c.遍城Biancheng
}
    
func (c *岳普湖YopurgaCounty) BFanjiaxiaozhang范家小掌() feud.Barony {
	return c.范家小掌Fanjiaxiaozhang
}
    
func (c *岳普湖YopurgaCounty) BHoushan后善() feud.Barony {
	return c.后善Houshan
}
    
func (c *岳普湖YopurgaCounty) BJiashi伽师() feud.Barony {
	return c.伽师Jiashi
}
    
func (c *岳普湖YopurgaCounty) BPalichuang八里庄() feud.Barony {
	return c.八里庄Palichuang
}
    
func (c *岳普湖YopurgaCounty) BYopurga岳普湖() feud.Barony {
	return c.岳普湖Yopurga
}
    
var CYopurga岳普湖 YopurgaCounty = &岳普湖YopurgaCounty{}

func init() {
	f := CYopurga岳普湖.(*岳普湖YopurgaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1525",
		Title:     "yopurga",
		TitleName: "岳普湖",
		TitleCode: "c_yopurga",
		Baronies:  map[string]feud.Barony{},
	}

	f.北降Beijiang = BBeijiang北降
	f.北降Beijiang.SetParent(f)
	
	f.遍城Biancheng = BBiancheng遍城
	f.遍城Biancheng.SetParent(f)
	
	f.范家小掌Fanjiaxiaozhang = BFanjiaxiaozhang范家小掌
	f.范家小掌Fanjiaxiaozhang.SetParent(f)
	
	f.后善Houshan = BHoushan后善
	f.后善Houshan.SetParent(f)
	
	f.伽师Jiashi = BJiashi伽师
	f.伽师Jiashi.SetParent(f)
	
	f.八里庄Palichuang = BPalichuang八里庄
	f.八里庄Palichuang.SetParent(f)
	
	f.岳普湖Yopurga = BYopurga岳普湖
	f.岳普湖Yopurga.SetParent(f)
	
}
