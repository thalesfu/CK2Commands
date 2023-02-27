package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UgraCounty interface {
    feud.County
    BKadzherom卡杰罗姆() 	feud.Barony
    BNizhodes下奥杰斯() 	feud.Barony
    BSosnogorsk索斯诺戈尔斯克() 	feud.Barony
    BUkhta乌赫塔() 	feud.Barony
    BVodnyy沃德内() 	feud.Barony
    BVoyvozh沃伊沃日() 	feud.Barony
    BVuktyl武克特尔() 	feud.Barony
    BYarega亚列加() 	feud.Barony
}

type 乌格拉UgraCounty struct {
	feud.BaseCounty
	卡杰罗姆Kadzherom 	feud.Barony
	下奥杰斯Nizhodes 	feud.Barony
	索斯诺戈尔斯克Sosnogorsk 	feud.Barony
	乌赫塔Ukhta 	feud.Barony
	沃德内Vodnyy 	feud.Barony
	沃伊沃日Voyvozh 	feud.Barony
	武克特尔Vuktyl 	feud.Barony
	亚列加Yarega 	feud.Barony
}

func (c *乌格拉UgraCounty) BKadzherom卡杰罗姆() feud.Barony {
	return c.卡杰罗姆Kadzherom
}
    
func (c *乌格拉UgraCounty) BNizhodes下奥杰斯() feud.Barony {
	return c.下奥杰斯Nizhodes
}
    
func (c *乌格拉UgraCounty) BSosnogorsk索斯诺戈尔斯克() feud.Barony {
	return c.索斯诺戈尔斯克Sosnogorsk
}
    
func (c *乌格拉UgraCounty) BUkhta乌赫塔() feud.Barony {
	return c.乌赫塔Ukhta
}
    
func (c *乌格拉UgraCounty) BVodnyy沃德内() feud.Barony {
	return c.沃德内Vodnyy
}
    
func (c *乌格拉UgraCounty) BVoyvozh沃伊沃日() feud.Barony {
	return c.沃伊沃日Voyvozh
}
    
func (c *乌格拉UgraCounty) BVuktyl武克特尔() feud.Barony {
	return c.武克特尔Vuktyl
}
    
func (c *乌格拉UgraCounty) BYarega亚列加() feud.Barony {
	return c.亚列加Yarega
}
    
var CUgra乌格拉 UgraCounty = &乌格拉UgraCounty{}

func init() {
	f := CUgra乌格拉.(*乌格拉UgraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "398",
		Title:     "ugra",
		TitleName: "乌格拉",
		TitleCode: "c_ugra",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡杰罗姆Kadzherom = BKadzherom卡杰罗姆
	f.卡杰罗姆Kadzherom.SetParent(f)
	
	f.下奥杰斯Nizhodes = BNizhodes下奥杰斯
	f.下奥杰斯Nizhodes.SetParent(f)
	
	f.索斯诺戈尔斯克Sosnogorsk = BSosnogorsk索斯诺戈尔斯克
	f.索斯诺戈尔斯克Sosnogorsk.SetParent(f)
	
	f.乌赫塔Ukhta = BUkhta乌赫塔
	f.乌赫塔Ukhta.SetParent(f)
	
	f.沃德内Vodnyy = BVodnyy沃德内
	f.沃德内Vodnyy.SetParent(f)
	
	f.沃伊沃日Voyvozh = BVoyvozh沃伊沃日
	f.沃伊沃日Voyvozh.SetParent(f)
	
	f.武克特尔Vuktyl = BVuktyl武克特尔
	f.武克特尔Vuktyl.SetParent(f)
	
	f.亚列加Yarega = BYarega亚列加
	f.亚列加Yarega.SetParent(f)
	
}
