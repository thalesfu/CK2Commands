package tarvagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TarvagataiCounty interface {
    feud.County
    BAsgat阿斯嘎特() 	feud.Barony
    BKungurtug昆古尔图格() 	feud.Barony
    BShuurmak舒尔马克() 	feud.Barony
    BSongino松吉瑙() 	feud.Barony
    BTarvagatai塔尔巴嘎泰() 	feud.Barony
    BTere_khol捷列霍尔() 	feud.Barony
    BZuunkhangai宗杭爱() 	feud.Barony
}

type 塔尔巴嘎泰TarvagataiCounty struct {
	feud.BaseCounty
	阿斯嘎特Asgat 	feud.Barony
	昆古尔图格Kungurtug 	feud.Barony
	舒尔马克Shuurmak 	feud.Barony
	松吉瑙Songino 	feud.Barony
	塔尔巴嘎泰Tarvagatai 	feud.Barony
	捷列霍尔Tere_khol 	feud.Barony
	宗杭爱Zuunkhangai 	feud.Barony
}

func (c *塔尔巴嘎泰TarvagataiCounty) BAsgat阿斯嘎特() feud.Barony {
	return c.阿斯嘎特Asgat
}
    
func (c *塔尔巴嘎泰TarvagataiCounty) BKungurtug昆古尔图格() feud.Barony {
	return c.昆古尔图格Kungurtug
}
    
func (c *塔尔巴嘎泰TarvagataiCounty) BShuurmak舒尔马克() feud.Barony {
	return c.舒尔马克Shuurmak
}
    
func (c *塔尔巴嘎泰TarvagataiCounty) BSongino松吉瑙() feud.Barony {
	return c.松吉瑙Songino
}
    
func (c *塔尔巴嘎泰TarvagataiCounty) BTarvagatai塔尔巴嘎泰() feud.Barony {
	return c.塔尔巴嘎泰Tarvagatai
}
    
func (c *塔尔巴嘎泰TarvagataiCounty) BTere_khol捷列霍尔() feud.Barony {
	return c.捷列霍尔Tere_khol
}
    
func (c *塔尔巴嘎泰TarvagataiCounty) BZuunkhangai宗杭爱() feud.Barony {
	return c.宗杭爱Zuunkhangai
}
    
var CTarvagatai塔尔巴嘎泰 TarvagataiCounty = &塔尔巴嘎泰TarvagataiCounty{}

func init() {
	f := CTarvagatai塔尔巴嘎泰.(*塔尔巴嘎泰TarvagataiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1914",
		Title:     "tarvagatai",
		TitleName: "塔尔巴嘎泰",
		TitleCode: "c_tarvagatai",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯嘎特Asgat = BAsgat阿斯嘎特
	f.阿斯嘎特Asgat.SetParent(f)
	
	f.昆古尔图格Kungurtug = BKungurtug昆古尔图格
	f.昆古尔图格Kungurtug.SetParent(f)
	
	f.舒尔马克Shuurmak = BShuurmak舒尔马克
	f.舒尔马克Shuurmak.SetParent(f)
	
	f.松吉瑙Songino = BSongino松吉瑙
	f.松吉瑙Songino.SetParent(f)
	
	f.塔尔巴嘎泰Tarvagatai = BTarvagatai塔尔巴嘎泰
	f.塔尔巴嘎泰Tarvagatai.SetParent(f)
	
	f.捷列霍尔Tere_khol = BTere_khol捷列霍尔
	f.捷列霍尔Tere_khol.SetParent(f)
	
	f.宗杭爱Zuunkhangai = BZuunkhangai宗杭爱
	f.宗杭爱Zuunkhangai.SetParent(f)
	
}
