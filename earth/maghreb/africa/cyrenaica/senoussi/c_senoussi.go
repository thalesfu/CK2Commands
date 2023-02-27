package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SenoussiCounty interface {
    feud.County
    BAttaj塔季() 	feud.Barony
    BAttallab泰拉卜() 	feud.Barony
    BBuzayman布宰迈() 	feud.Barony
    BGabr加布罗() 	feud.Barony
    BJakharrah贾哈尔拉() 	feud.Barony
    BJalu贾卢() 	feud.Barony
    BKufra库夫拉() 	feud.Barony
    BTazirbu塔济尔布() 	feud.Barony
}

type 塞努西SenoussiCounty struct {
	feud.BaseCounty
	塔季Attaj 	feud.Barony
	泰拉卜Attallab 	feud.Barony
	布宰迈Buzayman 	feud.Barony
	加布罗Gabr 	feud.Barony
	贾哈尔拉Jakharrah 	feud.Barony
	贾卢Jalu 	feud.Barony
	库夫拉Kufra 	feud.Barony
	塔济尔布Tazirbu 	feud.Barony
}

func (c *塞努西SenoussiCounty) BAttaj塔季() feud.Barony {
	return c.塔季Attaj
}
    
func (c *塞努西SenoussiCounty) BAttallab泰拉卜() feud.Barony {
	return c.泰拉卜Attallab
}
    
func (c *塞努西SenoussiCounty) BBuzayman布宰迈() feud.Barony {
	return c.布宰迈Buzayman
}
    
func (c *塞努西SenoussiCounty) BGabr加布罗() feud.Barony {
	return c.加布罗Gabr
}
    
func (c *塞努西SenoussiCounty) BJakharrah贾哈尔拉() feud.Barony {
	return c.贾哈尔拉Jakharrah
}
    
func (c *塞努西SenoussiCounty) BJalu贾卢() feud.Barony {
	return c.贾卢Jalu
}
    
func (c *塞努西SenoussiCounty) BKufra库夫拉() feud.Barony {
	return c.库夫拉Kufra
}
    
func (c *塞努西SenoussiCounty) BTazirbu塔济尔布() feud.Barony {
	return c.塔济尔布Tazirbu
}
    
var CSenoussi塞努西 SenoussiCounty = &塞努西SenoussiCounty{}

func init() {
	f := CSenoussi塞努西.(*塞努西SenoussiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "807",
		Title:     "senoussi",
		TitleName: "塞努西",
		TitleCode: "c_senoussi",
		Baronies:  map[string]feud.Barony{},
	}

	f.塔季Attaj = BAttaj塔季
	f.塔季Attaj.SetParent(f)
	
	f.泰拉卜Attallab = BAttallab泰拉卜
	f.泰拉卜Attallab.SetParent(f)
	
	f.布宰迈Buzayman = BBuzayman布宰迈
	f.布宰迈Buzayman.SetParent(f)
	
	f.加布罗Gabr = BGabr加布罗
	f.加布罗Gabr.SetParent(f)
	
	f.贾哈尔拉Jakharrah = BJakharrah贾哈尔拉
	f.贾哈尔拉Jakharrah.SetParent(f)
	
	f.贾卢Jalu = BJalu贾卢
	f.贾卢Jalu.SetParent(f)
	
	f.库夫拉Kufra = BKufra库夫拉
	f.库夫拉Kufra.SetParent(f)
	
	f.塔济尔布Tazirbu = BTazirbu塔济尔布
	f.塔济尔布Tazirbu.SetParent(f)
	
}
