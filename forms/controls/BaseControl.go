package controls

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
)

type BaseControl struct {
	FormBaseUI

	impl   p.Control
	isInit bool
	anchor f.AnchorStyle
}

func (_this *BaseControl) Init(impl p.Control) *BaseControl {
	_this.impl = impl
	_this.FormBaseUI.init(_this, _this.impl)
	_this.anchor = f.AnchorStyle_Left | f.AnchorStyle_Top
	_this.isInit = true
	return _this
}

func (_this *BaseControl) getAnchor() f.AnchorStyle {
	return _this.anchor
}

func (_this *BaseControl) SetAnchor(style f.AnchorStyle) {
	_this.anchor = style
}

func (_this *BaseControl) toControl() p.Control {
	return _this.impl
}

func (_this *BaseControl) Show() {
	_this.getImpl().Show()
}

func (_this *BaseControl) Hide() {
	_this.getImpl().Hide()
}

func (_this *BaseControl) getImpl() p.Control {
	if _this.isInit == false {
		panic("必须使用Init()初始化 ")
	}
	return _this.impl
}
