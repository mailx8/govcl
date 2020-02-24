
//----------------------------------------
// The code is automatically generated by the GenlibVcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TTaskDialogBaseButtonItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTaskDialogBaseButtonItem
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTaskDialogBaseButtonItem() *TTaskDialogBaseButtonItem {
    t := new(TTaskDialogBaseButtonItem)
    t.instance = TaskDialogBaseButtonItem_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TaskDialogBaseButtonItemFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TaskDialogBaseButtonItemFromInst(inst uintptr) *TTaskDialogBaseButtonItem {
    t := new(TTaskDialogBaseButtonItem)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TaskDialogBaseButtonItemFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TaskDialogBaseButtonItemFromObj(obj IObject) *TTaskDialogBaseButtonItem {
    t := new(TTaskDialogBaseButtonItem)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TaskDialogBaseButtonItemFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TaskDialogBaseButtonItemFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialogBaseButtonItem {
    t := new(TTaskDialogBaseButtonItem)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTaskDialogBaseButtonItem) Free() {
    if t.instance != 0 {
        TaskDialogBaseButtonItem_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskDialogBaseButtonItem) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskDialogBaseButtonItem) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskDialogBaseButtonItem) IsValid() bool {
    return t.instance != 0
}

// TTaskDialogBaseButtonItemClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskDialogBaseButtonItemClass() TClass {
    return TaskDialogBaseButtonItem_StaticClassType()
}

// Click
// CN: 单击。
// EN: .
func (t *TTaskDialogBaseButtonItem) Click() {
    TaskDialogBaseButtonItem_Click(t.instance)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskDialogBaseButtonItem) GetNamePath() string {
    return TaskDialogBaseButtonItem_GetNamePath(t.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskDialogBaseButtonItem) Assign(Source IObject) {
    TaskDialogBaseButtonItem_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTaskDialogBaseButtonItem) DisposeOf() {
    TaskDialogBaseButtonItem_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskDialogBaseButtonItem) ClassType() TClass {
    return TaskDialogBaseButtonItem_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskDialogBaseButtonItem) ClassName() string {
    return TaskDialogBaseButtonItem_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskDialogBaseButtonItem) InstanceSize() int32 {
    return TaskDialogBaseButtonItem_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskDialogBaseButtonItem) InheritsFrom(AClass TClass) bool {
    return TaskDialogBaseButtonItem_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskDialogBaseButtonItem) Equals(Obj IObject) bool {
    return TaskDialogBaseButtonItem_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskDialogBaseButtonItem) GetHashCode() int32 {
    return TaskDialogBaseButtonItem_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTaskDialogBaseButtonItem) ToString() string {
    return TaskDialogBaseButtonItem_ToString(t.instance)
}

// ModalResult
// CN: 获取模态对话框显示结果。
// EN: .
func (t *TTaskDialogBaseButtonItem) ModalResult() TModalResult {
    return TaskDialogBaseButtonItem_GetModalResult(t.instance)
}

// SetModalResult
// CN: 设置模态对话框显示结果。
// EN: .
func (t *TTaskDialogBaseButtonItem) SetModalResult(value TModalResult) {
    TaskDialogBaseButtonItem_SetModalResult(t.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (t *TTaskDialogBaseButtonItem) Caption() string {
    return TaskDialogBaseButtonItem_GetCaption(t.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (t *TTaskDialogBaseButtonItem) SetCaption(value string) {
    TaskDialogBaseButtonItem_SetCaption(t.instance, value)
}

// Default
func (t *TTaskDialogBaseButtonItem) Default() bool {
    return TaskDialogBaseButtonItem_GetDefault(t.instance)
}

// SetDefault
func (t *TTaskDialogBaseButtonItem) SetDefault(value bool) {
    TaskDialogBaseButtonItem_SetDefault(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTaskDialogBaseButtonItem) Enabled() bool {
    return TaskDialogBaseButtonItem_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTaskDialogBaseButtonItem) SetEnabled(value bool) {
    TaskDialogBaseButtonItem_SetEnabled(t.instance, value)
}

// Collection
func (t *TTaskDialogBaseButtonItem) Collection() *TCollection {
    return CollectionFromInst(TaskDialogBaseButtonItem_GetCollection(t.instance))
}

// SetCollection
func (t *TTaskDialogBaseButtonItem) SetCollection(value *TCollection) {
    TaskDialogBaseButtonItem_SetCollection(t.instance, CheckPtr(value))
}

// Index
func (t *TTaskDialogBaseButtonItem) Index() int32 {
    return TaskDialogBaseButtonItem_GetIndex(t.instance)
}

// SetIndex
func (t *TTaskDialogBaseButtonItem) SetIndex(value int32) {
    TaskDialogBaseButtonItem_SetIndex(t.instance, value)
}

// DisplayName
func (t *TTaskDialogBaseButtonItem) DisplayName() string {
    return TaskDialogBaseButtonItem_GetDisplayName(t.instance)
}

// SetDisplayName
func (t *TTaskDialogBaseButtonItem) SetDisplayName(value string) {
    TaskDialogBaseButtonItem_SetDisplayName(t.instance, value)
}

