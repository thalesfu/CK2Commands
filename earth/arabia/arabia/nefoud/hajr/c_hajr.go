package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HajrCounty interface {
    feud.County
    BAbha艾卜哈() 	feud.Barony
    BAlhudaydah荷台达() 	feud.Barony
    BAlmaqah阿尔玛达() 	feud.Barony
    BAssalif塞里夫() 	feud.Barony
    BBaljurashi拜勒朱尔希() 	feud.Barony
    BBisha比沙() 	feud.Barony
    BJizan吉赞() 	feud.Barony
    BKhamismushait海米斯穆谢特() 	feud.Barony
}

type 哈杰尔HajrCounty struct {
	feud.BaseCounty
	艾卜哈Abha 	feud.Barony
	荷台达Alhudaydah 	feud.Barony
	阿尔玛达Almaqah 	feud.Barony
	塞里夫Assalif 	feud.Barony
	拜勒朱尔希Baljurashi 	feud.Barony
	比沙Bisha 	feud.Barony
	吉赞Jizan 	feud.Barony
	海米斯穆谢特Khamismushait 	feud.Barony
}

func (c *哈杰尔HajrCounty) BAbha艾卜哈() feud.Barony {
	return c.艾卜哈Abha
}
    
func (c *哈杰尔HajrCounty) BAlhudaydah荷台达() feud.Barony {
	return c.荷台达Alhudaydah
}
    
func (c *哈杰尔HajrCounty) BAlmaqah阿尔玛达() feud.Barony {
	return c.阿尔玛达Almaqah
}
    
func (c *哈杰尔HajrCounty) BAssalif塞里夫() feud.Barony {
	return c.塞里夫Assalif
}
    
func (c *哈杰尔HajrCounty) BBaljurashi拜勒朱尔希() feud.Barony {
	return c.拜勒朱尔希Baljurashi
}
    
func (c *哈杰尔HajrCounty) BBisha比沙() feud.Barony {
	return c.比沙Bisha
}
    
func (c *哈杰尔HajrCounty) BJizan吉赞() feud.Barony {
	return c.吉赞Jizan
}
    
func (c *哈杰尔HajrCounty) BKhamismushait海米斯穆谢特() feud.Barony {
	return c.海米斯穆谢特Khamismushait
}
    
var CHajr哈杰尔 HajrCounty = &哈杰尔HajrCounty{}

func init() {
	f := CHajr哈杰尔.(*哈杰尔HajrCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "863",
		Title:     "hajr",
		TitleName: "哈杰尔",
		TitleCode: "c_hajr",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾卜哈Abha = BAbha艾卜哈
	f.艾卜哈Abha.SetParent(f)
	
	f.荷台达Alhudaydah = BAlhudaydah荷台达
	f.荷台达Alhudaydah.SetParent(f)
	
	f.阿尔玛达Almaqah = BAlmaqah阿尔玛达
	f.阿尔玛达Almaqah.SetParent(f)
	
	f.塞里夫Assalif = BAssalif塞里夫
	f.塞里夫Assalif.SetParent(f)
	
	f.拜勒朱尔希Baljurashi = BBaljurashi拜勒朱尔希
	f.拜勒朱尔希Baljurashi.SetParent(f)
	
	f.比沙Bisha = BBisha比沙
	f.比沙Bisha.SetParent(f)
	
	f.吉赞Jizan = BJizan吉赞
	f.吉赞Jizan.SetParent(f)
	
	f.海米斯穆谢特Khamismushait = BKhamismushait海米斯穆谢特
	f.海米斯穆谢特Khamismushait.SetParent(f)
	
}
