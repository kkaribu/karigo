package karigo

import (
	"time"
)

// Access ...
type Access struct {
	deps   []dep
	values map[Key]string
	ops    []Op
	done   bool
}

type dep struct {
	query *Query
	val   chan string
}

// Ready ...
// func (a *Access) Ready() {}

// Release ...
func (a Access) Release(query Query, keep ...[]string) {}

// WillGet ...
func (a *Access) WillGet(query Query) {}

// GetField ...
func (a *Access) GetField(key Key) string {
	a.deps = append(a.deps, dep{
		query: NewQueryFromKey(key),
		val:   make(chan string),
	})
	return ""
}

// GetString ...
func (a *Access) GetString(key Key) string { return "" }

// GetInt ...
func (a *Access) GetInt(key Key) int { return 0 }

// GetInt8 ...
func (a *Access) GetInt8(key Key) int8 { return 0 }

// GetInt16 ...
func (a *Access) GetInt16(key Key) int16 { return 0 }

// GetInt32 ...
func (a *Access) GetInt32(key Key) int32 { return 0 }

// GetInt64 ...
func (a *Access) GetInt64(key Key) int64 { return 0 }

// GetUint ...
func (a *Access) GetUint(key Key) uint { return 0 }

// GetUint8 ...
func (a *Access) GetUint8(key Key) uint8 { return 0 }

// GetUint16 ...
func (a *Access) GetUint16(key Key) uint16 { return 0 }

// GetUint32 ...
func (a *Access) GetUint32(key Key) uint32 { return 0 }

// GetUint64 ...
func (a *Access) GetUint64(key Key) uint64 { return 0 }

// GetBool ...
func (a *Access) GetBool(key Key) bool { return false }

// GetTime ...
func (a *Access) GetTime(key Key) time.Time { return time.Time{} }

// GetStringPtr ...
func (a *Access) GetStringPtr(key Key) *string { return nil }

// GetIntPtr ...
func (a *Access) GetIntPtr(key Key) *int { return nil }

// GetInt8Ptr ...
func (a *Access) GetInt8Ptr(key Key) *int8 { return nil }

// GetInt16Ptr ...
func (a *Access) GetInt16Ptr(key Key) *int16 { return nil }

// GetInt32Ptr ...
func (a *Access) GetInt32Ptr(key Key) *int32 { return nil }

// GetInt64Ptr ...
func (a *Access) GetInt64Ptr(key Key) *int64 { return nil }

// GetUintPtr ...
func (a *Access) GetUintPtr(key Key) *uint { return nil }

// GetUint8Ptr ...
func (a *Access) GetUint8Ptr(key Key) *uint8 { return nil }

// GetUint16Ptr ...
func (a *Access) GetUint16Ptr(key Key) *uint16 { return nil }

// GetUint32Ptr ...
func (a *Access) GetUint32Ptr(key Key) *uint32 { return nil }

// GetUint64Ptr ...
func (a *Access) GetUint64Ptr(key Key) *uint64 { return nil }

// GetBoolPtr ...
func (a *Access) GetBoolPtr(key Key) *bool { return nil }

// GetTimePtr ...
func (a *Access) GetTimePtr(key Key) *time.Time { return nil }

// GetStrings ...
func (a *Access) GetStrings(query Query) []string {
	return []string{}
}

// GetResFields ...
func (a *Access) GetResFields(query Query) map[string]interface{} {
	return map[string]interface{}{}
}

// GetInts ...
func (a *Access) GetInts(query Query) []int {
	return []int{}
}

// GetInt8s ...
func (a *Access) GetInt8s(query Query) []int8 {
	return []int8{}
}

// GetInt16s ...
func (a *Access) GetInt16s(query Query) []int16 {
	return []int16{}
}

// GetInt32s ...
func (a *Access) GetInt32s(query Query) []int32 {
	return []int32{}
}

// GetInt64s ...
func (a *Access) GetInt64s(query Query) []int64 {
	return []int64{}
}

// GetUints ...
func (a *Access) GetUints(query Query) []uint {
	return []uint{}
}

// GetUint8s ...
func (a *Access) GetUint8s(query Query) []uint8 {
	return []uint8{}
}

// GetUint16s ...
func (a *Access) GetUint16s(query Query) []uint16 {
	return []uint16{}
}

// GetUint32s ...
func (a *Access) GetUint32s(query Query) []uint32 {
	return []uint32{}
}

// GetUint64s ...
func (a *Access) GetUint64s(query Query) []uint64 {
	return []uint64{}
}

// GetBools ...
func (a *Access) GetBools(query Query) []bool {
	return []bool{}
}

// GetTimes ...
func (a *Access) GetTimes(query Query) []time.Time {
	return []time.Time{}
}

// GetStringPtrs ...
func (a *Access) GetStringPtrs(query Query) []*string {
	return []*string{}
}

// GetIntPtrs ...
func (a *Access) GetIntPtrs(query Query) []*int {
	return []*int{}
}

// GetInt8Ptrs ...
func (a *Access) GetInt8Ptrs(query Query) []*int8 {
	return []*int8{}
}

// GetInt16Ptrs ...
func (a *Access) GetInt16Ptrs(query Query) []*int16 {
	return []*int16{}
}

// GetInt32Ptrs ...
func (a *Access) GetInt32Ptrs(query Query) []*int32 {
	return []*int32{}
}

// GetInt64Ptrs ...
func (a *Access) GetInt64Ptrs(query Query) []*int64 {
	return []*int64{}
}

// GetUintPtrs ...
func (a *Access) GetUintPtrs(query Query) []*uint {
	return []*uint{}
}

// GetUint8Ptrs ...
func (a *Access) GetUint8Ptrs(query Query) []*uint8 {
	return []*uint8{}
}

// GetUint16Ptrs ...
func (a *Access) GetUint16Ptrs(query Query) []*uint16 {
	return []*uint16{}
}

// GetUint32Ptrs ...
func (a *Access) GetUint32Ptrs(query Query) []*uint32 {
	return []*uint32{}
}

// GetUint64Ptrs ...
func (a *Access) GetUint64Ptrs(query Query) []*uint64 {
	return []*uint64{}
}

// GetBoolPtrs ...
func (a *Access) GetBoolPtrs(query Query) []*bool {
	return []*bool{}
}

// GetTimePtrs ...
func (a *Access) GetTimePtrs(query Query) []*time.Time {
	return []*time.Time{}
}

// GetColFields ...
func (a *Access) GetColFields(query Query) []map[string]interface{} {
	return []map[string]interface{}{}
}

// GetInclusions ...
func (a *Access) GetInclusions(query Query, rels, fields []string) map[string]map[string]interface{} {
	return map[string]map[string]interface{}{}
}

// GetSlice ...
// TODO What is that
// func (a *Access) GetSlice(query Query) []interface{} {
// 	return []interface{}{}
// }

// Count ...
func (a *Access) Count(query Query) int {
	return 0
}

// WillSet ...
func (a *Access) WillSet(query Query) {}

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

// set ...
func (a *Access) set(key Key, val interface{}) error { return nil }
