package barkul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BarkulCounty interface {
    feud.County
    BBarkul婆悉厥() 	feud.Barony
    BJiuxian旧县镇() 	feud.Barony
    BKhoid辉特() 	feud.Barony
    BShiniu石纽() 	feud.Barony
    BTiantoucun天头村() 	feud.Barony
    BXiaoguantian小关田() 	feud.Barony
    BZhangpeng漳澎() 	feud.Barony
}

type 婆悉厥BarkulCounty struct {
	feud.BaseCounty
	婆悉厥Barkul 	feud.Barony
	旧县镇Jiuxian 	feud.Barony
	辉特Khoid 	feud.Barony
	石纽Shiniu 	feud.Barony
	天头村Tiantoucun 	feud.Barony
	小关田Xiaoguantian 	feud.Barony
	漳澎Zhangpeng 	feud.Barony
}

func (c *婆悉厥BarkulCounty) BBarkul婆悉厥() feud.Barony {
	return c.婆悉厥Barkul
}
    
func (c *婆悉厥BarkulCounty) BJiuxian旧县镇() feud.Barony {
	return c.旧县镇Jiuxian
}
    
func (c *婆悉厥BarkulCounty) BKhoid辉特() feud.Barony {
	return c.辉特Khoid
}
    
func (c *婆悉厥BarkulCounty) BShiniu石纽() feud.Barony {
	return c.石纽Shiniu
}
    
func (c *婆悉厥BarkulCounty) BTiantoucun天头村() feud.Barony {
	return c.天头村Tiantoucun
}
    
func (c *婆悉厥BarkulCounty) BXiaoguantian小关田() feud.Barony {
	return c.小关田Xiaoguantian
}
    
func (c *婆悉厥BarkulCounty) BZhangpeng漳澎() feud.Barony {
	return c.漳澎Zhangpeng
}
    
var CBarkul婆悉厥 BarkulCounty = &婆悉厥BarkulCounty{}

func init() {
	f := CBarkul婆悉厥.(*婆悉厥BarkulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1516",
		Title:     "barkul",
		TitleName: "婆悉厥",
		TitleCode: "c_barkul",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆悉厥Barkul = BBarkul婆悉厥
	f.婆悉厥Barkul.SetParent(f)
	
	f.旧县镇Jiuxian = BJiuxian旧县镇
	f.旧县镇Jiuxian.SetParent(f)
	
	f.辉特Khoid = BKhoid辉特
	f.辉特Khoid.SetParent(f)
	
	f.石纽Shiniu = BShiniu石纽
	f.石纽Shiniu.SetParent(f)
	
	f.天头村Tiantoucun = BTiantoucun天头村
	f.天头村Tiantoucun.SetParent(f)
	
	f.小关田Xiaoguantian = BXiaoguantian小关田
	f.小关田Xiaoguantian.SetParent(f)
	
	f.漳澎Zhangpeng = BZhangpeng漳澎
	f.漳澎Zhangpeng.SetParent(f)
	
}
