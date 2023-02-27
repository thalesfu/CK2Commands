package altay

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/altay/belukha"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/altay/khuiten"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/altay/markakol"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/altay/monkh_khairkhan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/altay/muztau"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AltayDuke interface {
    feud.Duke
    CBelukha别卢哈() 	belukha.BelukhaCounty
    CKhuiten辉腾() 	khuiten.KhuitenCounty
    CMarkakol马尔卡科尔() 	markakol.MarkakolCounty
    CMonkh_khairkhan蒙赫海尔汗() 	monkh_khairkhan.Monkh_khairkhanCounty
    CMuztau西萨彦() 	muztau.MuztauCounty
}

type 阿尔泰AltayDuke struct {
	feud.BaseDuke
	别卢哈Belukha 	belukha.BelukhaCounty
	辉腾Khuiten 	khuiten.KhuitenCounty
	马尔卡科尔Markakol 	markakol.MarkakolCounty
	蒙赫海尔汗Monkh_khairkhan 	monkh_khairkhan.Monkh_khairkhanCounty
	西萨彦Muztau 	muztau.MuztauCounty
}

func (d *阿尔泰AltayDuke) CBelukha别卢哈() belukha.BelukhaCounty {
	return d.别卢哈Belukha
}
    
func (d *阿尔泰AltayDuke) CKhuiten辉腾() khuiten.KhuitenCounty {
	return d.辉腾Khuiten
}
    
func (d *阿尔泰AltayDuke) CMarkakol马尔卡科尔() markakol.MarkakolCounty {
	return d.马尔卡科尔Markakol
}
    
func (d *阿尔泰AltayDuke) CMonkh_khairkhan蒙赫海尔汗() monkh_khairkhan.Monkh_khairkhanCounty {
	return d.蒙赫海尔汗Monkh_khairkhan
}
    
func (d *阿尔泰AltayDuke) CMuztau西萨彦() muztau.MuztauCounty {
	return d.西萨彦Muztau
}
    
var DAltay阿尔泰 AltayDuke = &阿尔泰AltayDuke{}

func init() {
	f := DAltay阿尔泰.(*阿尔泰AltayDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "altay",
		TitleName: "阿尔泰",
		TitleCode: "d_altay",
		Counties:  map[string]feud.County{},
	}

	f.别卢哈Belukha = belukha.CBelukha别卢哈
	f.别卢哈Belukha.SetParent(f)
	
	f.辉腾Khuiten = khuiten.CKhuiten辉腾
	f.辉腾Khuiten.SetParent(f)
	
	f.马尔卡科尔Markakol = markakol.CMarkakol马尔卡科尔
	f.马尔卡科尔Markakol.SetParent(f)
	
	f.蒙赫海尔汗Monkh_khairkhan = monkh_khairkhan.CMonkh_khairkhan蒙赫海尔汗
	f.蒙赫海尔汗Monkh_khairkhan.SetParent(f)
	
	f.西萨彦Muztau = muztau.CMuztau西萨彦
	f.西萨彦Muztau.SetParent(f)
	
}
