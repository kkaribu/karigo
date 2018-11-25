package karigo

import (
	"time"
)

// Access ...
type Access struct {
	udpates map[SimpleKey]interface{}
}

// Ready ...
func (a *Access) Ready() {}

// Release ...
func (a Access) Release(key Key, keep ...[]string) {}

// WillGet ...
func (a *Access) WillGet(key Key) {}

// Get ...
func (a *Access) Get(key Key) interface{} { return nil }

// GetString ...
func (a *Access) GetString(key SimpleKey) string { return "" }

// GetInt ...
func (a *Access) GetInt(key SimpleKey) int { return 0 }

// GetInt8 ...
func (a *Access) GetInt8(key SimpleKey) int8 { return 0 }

// GetInt16 ...
func (a *Access) GetInt16(key SimpleKey) int16 { return 0 }

// GetInt32 ...
func (a *Access) GetInt32(key SimpleKey) int32 { return 0 }

// GetInt64 ...
func (a *Access) GetInt64(key SimpleKey) int64 { return 0 }

// GetUint ...
func (a *Access) GetUint(key SimpleKey) uint { return 0 }

// GetUint8 ...
func (a *Access) GetUint8(key SimpleKey) uint8 { return 0 }

// GetUint16 ...
func (a *Access) GetUint16(key SimpleKey) uint16 { return 0 }

// GetUint32 ...
func (a *Access) GetUint32(key SimpleKey) uint32 { return 0 }

// GetUint64 ...
func (a *Access) GetUint64(key SimpleKey) uint64 { return 0 }

// GetBool ...
func (a *Access) GetBool(key SimpleKey) bool { return false }

// GetTime ...
func (a *Access) GetTime(key SimpleKey) time.Time { return time.Time{} }

// GetStringPtr ...
func (a *Access) GetStringPtr(key SimpleKey) *string { return nil }

// GetIntPtr ...
func (a *Access) GetIntPtr(key SimpleKey) *int { return nil }

// GetInt8Ptr ...
func (a *Access) GetInt8Ptr(key SimpleKey) *int8 { return nil }

// GetInt16Ptr ...
func (a *Access) GetInt16Ptr(key SimpleKey) *int16 { return nil }

// GetInt32Ptr ...
func (a *Access) GetInt32Ptr(key SimpleKey) *int32 { return nil }

// GetInt64Ptr ...
func (a *Access) GetInt64Ptr(key SimpleKey) *int64 { return nil }

// GetUintPtr ...
func (a *Access) GetUintPtr(key SimpleKey) *uint { return nil }

// GetUint8Ptr ...
func (a *Access) GetUint8Ptr(key SimpleKey) *uint8 { return nil }

// GetUint16Ptr ...
func (a *Access) GetUint16Ptr(key SimpleKey) *uint16 { return nil }

// GetUint32Ptr ...
func (a *Access) GetUint32Ptr(key SimpleKey) *uint32 { return nil }

// GetUint64Ptr ...
func (a *Access) GetUint64Ptr(key SimpleKey) *uint64 { return nil }

// GetBoolPtr ...
func (a *Access) GetBoolPtr(key SimpleKey) *bool { return nil }

// GetTimePtr ...
func (a *Access) GetTimePtr(key SimpleKey) *time.Time { return nil }

// GetResFields ...
func (a *Access) GetResFields(key Key) map[string]interface{} {
	return map[string]interface{}{}
}

// GetSlice ...
func (a *Access) GetSlice(key Key) []interface{} {
	return []interface{}{}
}

// GetStrings ...
func (a *Access) GetStrings(key Key) []string {
	return []string{}
}

// GetInts ...
func (a *Access) GetInts(key Key) []int {
	return []int{}
}

// GetInt8s ...
func (a *Access) GetInt8s(key Key) []int8 {
	return []int8{}
}

// GetInt16s ...
func (a *Access) GetInt16s(key Key) []int16 {
	return []int16{}
}

// GetInt32s ...
func (a *Access) GetInt32s(key Key) []int32 {
	return []int32{}
}

// GetInt64s ...
func (a *Access) GetInt64s(key Key) []int64 {
	return []int64{}
}

// GetUints ...
func (a *Access) GetUints(key Key) []uint {
	return []uint{}
}

// GetUint8s ...
func (a *Access) GetUint8s(key Key) []uint8 {
	return []uint8{}
}

// GetUint16s ...
func (a *Access) GetUint16s(key Key) []uint16 {
	return []uint16{}
}

// GetUint32s ...
func (a *Access) GetUint32s(key Key) []uint32 {
	return []uint32{}
}

// GetUint64s ...
func (a *Access) GetUint64s(key Key) []uint64 {
	return []uint64{}
}

// GetBools ...
func (a *Access) GetBools(key Key) []bool {
	return []bool{}
}

// GetTimes ...
func (a *Access) GetTimes(key Key) []time.Time {
	return []time.Time{}
}

// GetStringPtrs ...
func (a *Access) GetStringPtrs(key Key) []*string {
	return []*string{}
}

// GetIntPtrs ...
func (a *Access) GetIntPtrs(key Key) []*int {
	return []*int{}
}

// GetInt8Ptrs ...
func (a *Access) GetInt8Ptrs(key Key) []*int8 {
	return []*int8{}
}

// GetInt16Ptrs ...
func (a *Access) GetInt16Ptrs(key Key) []*int16 {
	return []*int16{}
}

// GetInt32Ptrs ...
func (a *Access) GetInt32Ptrs(key Key) []*int32 {
	return []*int32{}
}

// GetInt64Ptrs ...
func (a *Access) GetInt64Ptrs(key Key) []*int64 {
	return []*int64{}
}

// GetUintPtrs ...
func (a *Access) GetUintPtrs(key Key) []*uint {
	return []*uint{}
}

// GetUint8Ptrs ...
func (a *Access) GetUint8Ptrs(key Key) []*uint8 {
	return []*uint8{}
}

// GetUint16Ptrs ...
func (a *Access) GetUint16Ptrs(key Key) []*uint16 {
	return []*uint16{}
}

// GetUint32Ptrs ...
func (a *Access) GetUint32Ptrs(key Key) []*uint32 {
	return []*uint32{}
}

// GetUint64Ptrs ...
func (a *Access) GetUint64Ptrs(key Key) []*uint64 {
	return []*uint64{}
}

// GetBoolPtrs ...
func (a *Access) GetBoolPtrs(key Key) []*bool {
	return []*bool{}
}

// GetTimePtrs ...
func (a *Access) GetTimePtrs(key Key) []*time.Time {
	return []*time.Time{}
}

// GetColFields ...
func (a *Access) GetColFields(key Key) []map[string]interface{} {
	return []map[string]interface{}{}
}

// WillSet ...
func (a *Access) WillSet(key Key) {}

// Set ...
func (a *Access) Set(typ, id, field string, v interface{}) {}

// SetString ...
func (a *Access) SetString(typ, id, field string, v string) {}

// SetInt ...
func (a *Access) SetInt(typ, id, field string, v int) {}

// SetInt8 ...
func (a *Access) SetInt8(typ, id, field string, v int8) {}

// SetInt16 ...
func (a *Access) SetInt16(typ, id, field string, v int16) {}

// SetInt32 ...
func (a *Access) SetInt32(typ, id, field string, v int32) {}

// SetInt64 ...
func (a *Access) SetInt64(typ, id, field string, v int64) {}

// SetUint ...
func (a *Access) SetUint(typ, id, field string, v uint) {}

// SetUint8 ...
func (a *Access) SetUint8(typ, id, field string, v uint8) {}

// SetUint16 ...
func (a *Access) SetUint16(typ, id, field string, v uint16) {}

// SetUint32 ...
func (a *Access) SetUint32(typ, id, field string, v uint32) {}

// SetUint64 ...
func (a *Access) SetUint64(typ, id, field string, v uint64) {}

// SetBool ...
func (a *Access) SetBool(typ, id, field string, v bool) {}

// SetTime ...
func (a *Access) SetTime(typ, id, field string, v time.Time) {}

// SetStringPtr ...
func (a *Access) SetStringPtr(typ, id, field string, v *string) {}

// SetIntPtr ...
func (a *Access) SetIntPtr(typ, id, field string, v *int) {}

// SetInt8Ptr ...
func (a *Access) SetInt8Ptr(typ, id, field string, v *int8) {}

// SetInt16Ptr ...
func (a *Access) SetInt16Ptr(typ, id, field string, v *int16) {}

// SetInt32Ptr ...
func (a *Access) SetInt32Ptr(typ, id, field string, v *int32) {}

// SetInt64Ptr ...
func (a *Access) SetInt64Ptr(typ, id, field string, v *int64) {}

// SetUintPtr ...
func (a *Access) SetUintPtr(typ, id, field string, v *uint) {}

// SetUint8Ptr ...
func (a *Access) SetUint8Ptr(typ, id, field string, v *uint8) {}

// SetUint16Ptr ...
func (a *Access) SetUint16Ptr(typ, id, field string, v *uint16) {}

// SetUint32Ptr ...
func (a *Access) SetUint32Ptr(typ, id, field string, v *uint32) {}

// SetUint64Ptr ...
func (a *Access) SetUint64Ptr(typ, id, field string, v *uint64) {}

// SetBoolPtr ...
func (a *Access) SetBoolPtr(typ, id, field string, v *bool) {}

// SetTimePtr ...
func (a *Access) SetTimePtr(typ, id, field string, v *time.Time) {}

// GetToOneRel ...
func (a *Access) GetToOneRel(key string) string { return "" }

// SetToOneRel ...
func (a *Access) SetToOneRel(typ, id, field string, nid string) {}

// GetToManyRel ...
func (a *Access) GetToManyRel(key string) []string { return []string{} }

// AddToManyRel ...
func (a *Access) AddToManyRel(key string, ids ...string) {}

// SetToManyRel ...
func (a *Access) SetToManyRel(typ, id, field string, ids ...string) {}

// DeleteToManyRel ...
func (a *Access) DeleteToManyRel(key string, ids ...string) {}
