package ivrea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IvreaCounty interface {
    feud.County
    BBard巴德() 	feud.Barony
    BBiella比耶拉() 	feud.Barony
    BCastellengo卡斯泰伦戈() 	feud.Barony
    BChamporcher尚波尔谢() 	feud.Barony
    BCossato科萨托() 	feud.Barony
    BIvrea伊夫雷亚() 	feud.Barony
    BPont_saint_martin蓬圣马丁() 	feud.Barony
}

type 伊夫雷亚IvreaCounty struct {
	feud.BaseCounty
	巴德Bard 	feud.Barony
	比耶拉Biella 	feud.Barony
	卡斯泰伦戈Castellengo 	feud.Barony
	尚波尔谢Champorcher 	feud.Barony
	科萨托Cossato 	feud.Barony
	伊夫雷亚Ivrea 	feud.Barony
	蓬圣马丁Pont_saint_martin 	feud.Barony
}

func (c *伊夫雷亚IvreaCounty) BBard巴德() feud.Barony {
	return c.巴德Bard
}
    
func (c *伊夫雷亚IvreaCounty) BBiella比耶拉() feud.Barony {
	return c.比耶拉Biella
}
    
func (c *伊夫雷亚IvreaCounty) BCastellengo卡斯泰伦戈() feud.Barony {
	return c.卡斯泰伦戈Castellengo
}
    
func (c *伊夫雷亚IvreaCounty) BChamporcher尚波尔谢() feud.Barony {
	return c.尚波尔谢Champorcher
}
    
func (c *伊夫雷亚IvreaCounty) BCossato科萨托() feud.Barony {
	return c.科萨托Cossato
}
    
func (c *伊夫雷亚IvreaCounty) BIvrea伊夫雷亚() feud.Barony {
	return c.伊夫雷亚Ivrea
}
    
func (c *伊夫雷亚IvreaCounty) BPont_saint_martin蓬圣马丁() feud.Barony {
	return c.蓬圣马丁Pont_saint_martin
}
    
var CIvrea伊夫雷亚 IvreaCounty = &伊夫雷亚IvreaCounty{}

func init() {
	f := CIvrea伊夫雷亚.(*伊夫雷亚IvreaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1710",
		Title:     "ivrea",
		TitleName: "伊夫雷亚",
		TitleCode: "c_ivrea",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴德Bard = BBard巴德
	f.巴德Bard.SetParent(f)
	
	f.比耶拉Biella = BBiella比耶拉
	f.比耶拉Biella.SetParent(f)
	
	f.卡斯泰伦戈Castellengo = BCastellengo卡斯泰伦戈
	f.卡斯泰伦戈Castellengo.SetParent(f)
	
	f.尚波尔谢Champorcher = BChamporcher尚波尔谢
	f.尚波尔谢Champorcher.SetParent(f)
	
	f.科萨托Cossato = BCossato科萨托
	f.科萨托Cossato.SetParent(f)
	
	f.伊夫雷亚Ivrea = BIvrea伊夫雷亚
	f.伊夫雷亚Ivrea.SetParent(f)
	
	f.蓬圣马丁Pont_saint_martin = BPont_saint_martin蓬圣马丁
	f.蓬圣马丁Pont_saint_martin.SetParent(f)
	
}
