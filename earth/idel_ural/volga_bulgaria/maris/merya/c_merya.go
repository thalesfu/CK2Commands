package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeryaCounty interface {
    feud.County
    BKilemary基列马雷() 	feud.Barony
    BLopatino洛帕季诺() 	feud.Barony
    BMariturek马里_图列克() 	feud.Barony
    BMorki莫尔基() 	feud.Barony
    BProvoi普罗沃伊() 	feud.Barony
    BTsikma齐克迈() 	feud.Barony
    BVolzhsk沃尔日斯克() 	feud.Barony
    BYoshkarola约什卡尔_奥拉() 	feud.Barony
}

type 马里MeryaCounty struct {
	feud.BaseCounty
	基列马雷Kilemary 	feud.Barony
	洛帕季诺Lopatino 	feud.Barony
	马里_图列克Mariturek 	feud.Barony
	莫尔基Morki 	feud.Barony
	普罗沃伊Provoi 	feud.Barony
	齐克迈Tsikma 	feud.Barony
	沃尔日斯克Volzhsk 	feud.Barony
	约什卡尔_奥拉Yoshkarola 	feud.Barony
}

func (c *马里MeryaCounty) BKilemary基列马雷() feud.Barony {
	return c.基列马雷Kilemary
}
    
func (c *马里MeryaCounty) BLopatino洛帕季诺() feud.Barony {
	return c.洛帕季诺Lopatino
}
    
func (c *马里MeryaCounty) BMariturek马里_图列克() feud.Barony {
	return c.马里_图列克Mariturek
}
    
func (c *马里MeryaCounty) BMorki莫尔基() feud.Barony {
	return c.莫尔基Morki
}
    
func (c *马里MeryaCounty) BProvoi普罗沃伊() feud.Barony {
	return c.普罗沃伊Provoi
}
    
func (c *马里MeryaCounty) BTsikma齐克迈() feud.Barony {
	return c.齐克迈Tsikma
}
    
func (c *马里MeryaCounty) BVolzhsk沃尔日斯克() feud.Barony {
	return c.沃尔日斯克Volzhsk
}
    
func (c *马里MeryaCounty) BYoshkarola约什卡尔_奥拉() feud.Barony {
	return c.约什卡尔_奥拉Yoshkarola
}
    
var CMerya马里 MeryaCounty = &马里MeryaCounty{}

func init() {
	f := CMerya马里.(*马里MeryaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "588",
		Title:     "merya",
		TitleName: "马里",
		TitleCode: "c_merya",
		Baronies:  map[string]feud.Barony{},
	}

	f.基列马雷Kilemary = BKilemary基列马雷
	f.基列马雷Kilemary.SetParent(f)
	
	f.洛帕季诺Lopatino = BLopatino洛帕季诺
	f.洛帕季诺Lopatino.SetParent(f)
	
	f.马里_图列克Mariturek = BMariturek马里_图列克
	f.马里_图列克Mariturek.SetParent(f)
	
	f.莫尔基Morki = BMorki莫尔基
	f.莫尔基Morki.SetParent(f)
	
	f.普罗沃伊Provoi = BProvoi普罗沃伊
	f.普罗沃伊Provoi.SetParent(f)
	
	f.齐克迈Tsikma = BTsikma齐克迈
	f.齐克迈Tsikma.SetParent(f)
	
	f.沃尔日斯克Volzhsk = BVolzhsk沃尔日斯克
	f.沃尔日斯克Volzhsk.SetParent(f)
	
	f.约什卡尔_奥拉Yoshkarola = BYoshkarola约什卡尔_奥拉
	f.约什卡尔_奥拉Yoshkarola.SetParent(f)
	
}
