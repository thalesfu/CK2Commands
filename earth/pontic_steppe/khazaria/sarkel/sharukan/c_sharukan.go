package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SharukanCounty interface {
    feud.County
    BBalakliia巴拉克列亚() 	feud.Barony
    BChallykala恰雷_卡拉() 	feud.Barony
    BIzyum伊久姆() 	feud.Barony
    BKharka哈尔卡() 	feud.Barony
    BKhorysdan霍雷斯丹() 	feud.Barony
    BKupyansk库普扬斯克() 	feud.Barony
    BLyubotin柳博京() 	feud.Barony
    BSumy苏梅() 	feud.Barony
}

type 沙鲁坎SharukanCounty struct {
	feud.BaseCounty
	巴拉克列亚Balakliia 	feud.Barony
	恰雷_卡拉Challykala 	feud.Barony
	伊久姆Izyum 	feud.Barony
	哈尔卡Kharka 	feud.Barony
	霍雷斯丹Khorysdan 	feud.Barony
	库普扬斯克Kupyansk 	feud.Barony
	柳博京Lyubotin 	feud.Barony
	苏梅Sumy 	feud.Barony
}

func (c *沙鲁坎SharukanCounty) BBalakliia巴拉克列亚() feud.Barony {
	return c.巴拉克列亚Balakliia
}
    
func (c *沙鲁坎SharukanCounty) BChallykala恰雷_卡拉() feud.Barony {
	return c.恰雷_卡拉Challykala
}
    
func (c *沙鲁坎SharukanCounty) BIzyum伊久姆() feud.Barony {
	return c.伊久姆Izyum
}
    
func (c *沙鲁坎SharukanCounty) BKharka哈尔卡() feud.Barony {
	return c.哈尔卡Kharka
}
    
func (c *沙鲁坎SharukanCounty) BKhorysdan霍雷斯丹() feud.Barony {
	return c.霍雷斯丹Khorysdan
}
    
func (c *沙鲁坎SharukanCounty) BKupyansk库普扬斯克() feud.Barony {
	return c.库普扬斯克Kupyansk
}
    
func (c *沙鲁坎SharukanCounty) BLyubotin柳博京() feud.Barony {
	return c.柳博京Lyubotin
}
    
func (c *沙鲁坎SharukanCounty) BSumy苏梅() feud.Barony {
	return c.苏梅Sumy
}
    
var CSharukan沙鲁坎 SharukanCounty = &沙鲁坎SharukanCounty{}

func init() {
	f := CSharukan沙鲁坎.(*沙鲁坎SharukanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "565",
		Title:     "sharukan",
		TitleName: "沙鲁坎",
		TitleCode: "c_sharukan",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴拉克列亚Balakliia = BBalakliia巴拉克列亚
	f.巴拉克列亚Balakliia.SetParent(f)
	
	f.恰雷_卡拉Challykala = BChallykala恰雷_卡拉
	f.恰雷_卡拉Challykala.SetParent(f)
	
	f.伊久姆Izyum = BIzyum伊久姆
	f.伊久姆Izyum.SetParent(f)
	
	f.哈尔卡Kharka = BKharka哈尔卡
	f.哈尔卡Kharka.SetParent(f)
	
	f.霍雷斯丹Khorysdan = BKhorysdan霍雷斯丹
	f.霍雷斯丹Khorysdan.SetParent(f)
	
	f.库普扬斯克Kupyansk = BKupyansk库普扬斯克
	f.库普扬斯克Kupyansk.SetParent(f)
	
	f.柳博京Lyubotin = BLyubotin柳博京
	f.柳博京Lyubotin.SetParent(f)
	
	f.苏梅Sumy = BSumy苏梅
	f.苏梅Sumy.SetParent(f)
	
}
