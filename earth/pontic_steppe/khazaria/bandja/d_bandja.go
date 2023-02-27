package bandja

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/bandja/lower_volga"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/bandja/pecheneg"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/bandja/uvek"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BandjaDuke interface {
    feud.Duke
    CLower_volga下伏尔加() 	lower_volga.Lower_volgaCounty
    CPecheneg班贾() 	pecheneg.PechenegCounty
    CUvek乌韦克() 	uvek.UvekCounty
}

type 班贾BandjaDuke struct {
	feud.BaseDuke
	下伏尔加Lower_volga 	lower_volga.Lower_volgaCounty
	班贾Pecheneg 	pecheneg.PechenegCounty
	乌韦克Uvek 	uvek.UvekCounty
}

func (d *班贾BandjaDuke) CLower_volga下伏尔加() lower_volga.Lower_volgaCounty {
	return d.下伏尔加Lower_volga
}
    
func (d *班贾BandjaDuke) CPecheneg班贾() pecheneg.PechenegCounty {
	return d.班贾Pecheneg
}
    
func (d *班贾BandjaDuke) CUvek乌韦克() uvek.UvekCounty {
	return d.乌韦克Uvek
}
    
var DBandja班贾 BandjaDuke = &班贾BandjaDuke{}

func init() {
	f := DBandja班贾.(*班贾BandjaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bandja",
		TitleName: "班贾",
		TitleCode: "d_bandja",
		Counties:  map[string]feud.County{},
	}

	f.下伏尔加Lower_volga = lower_volga.CLower_volga下伏尔加
	f.下伏尔加Lower_volga.SetParent(f)
	
	f.班贾Pecheneg = pecheneg.CPecheneg班贾
	f.班贾Pecheneg.SetParent(f)
	
	f.乌韦克Uvek = uvek.CUvek乌韦克
	f.乌韦克Uvek.SetParent(f)
	
}
