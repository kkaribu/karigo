package karigo

import (
	"time"

	"github.com/kkaribu/jsonapi"
)

// Access ...
type Access struct{}

// Ready ...
func (a *Access) Ready() {}

// WillGet ...
func (a *Access) WillGet(lock string) {}

// GetString ...
func (a *Access) GetString(lock string) string { return "" }

// GetInt ...
func (a *Access) GetInt(lock string) int { return 0 }

// GetInt8 ...
func (a *Access) GetInt8(lock string) int8 { return 0 }

// GetInt16 ...
func (a *Access) GetInt16(lock string) int16 { return 0 }

// GetInt32 ...
func (a *Access) GetInt32(lock string) int32 { return 0 }

// GetInt64 ...
func (a *Access) GetInt64(lock string) int64 { return 0 }

// GetUint ...
func (a *Access) GetUint(lock string) uint { return 0 }

// GetUint8 ...
func (a *Access) GetUint8(lock string) uint8 { return 0 }

// GetUint16 ...
func (a *Access) GetUint16(lock string) uint16 { return 0 }

// GetUint32 ...
func (a *Access) GetUint32(lock string) uint32 { return 0 }

// GetUint64 ...
func (a *Access) GetUint64(lock string) uint64 { return 0 }

// GetBool ...
func (a *Access) GetBool(lock string) bool { return false }

// GetTime ...
func (a *Access) GetTime(lock string) time.Time { return time.Time{} }

// GetStringPtr ...
func (a *Access) GetStringPtr(lock string) *string { return nil }

// GetIntPtr ...
func (a *Access) GetIntPtr(lock string) *int { return nil }

// GetInt8Ptr ...
func (a *Access) GetInt8Ptr(lock string) *int8 { return nil }

// GetInt16Ptr ...
func (a *Access) GetInt16Ptr(lock string) *int16 { return nil }

// GetInt32Ptr ...
func (a *Access) GetInt32Ptr(lock string) *int32 { return nil }

// GetInt64Ptr ...
func (a *Access) GetInt64Ptr(lock string) *int64 { return nil }

// GetUintPtr ...
func (a *Access) GetUintPtr(lock string) *uint { return nil }

// GetUint8Ptr ...
func (a *Access) GetUint8Ptr(lock string) *uint8 { return nil }

// GetUint16Ptr ...
func (a *Access) GetUint16Ptr(lock string) *uint16 { return nil }

// GetUint32Ptr ...
func (a *Access) GetUint32Ptr(lock string) *uint32 { return nil }

// GetUint64Ptr ...
func (a *Access) GetUint64Ptr(lock string) *uint64 { return nil }

// GetBoolPtr ...
func (a *Access) GetBoolPtr(lock string) *bool { return nil }

// GetTimePtr ...
func (a *Access) GetTimePtr(lock string) *time.Time { return nil }

// GetStrings ...
func (a *Access) GetStrings(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []string {
	return []string{}
}

// GetInts ...
func (a *Access) GetInts(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []int {
	return []int{}
}

// GetInt8s ...
func (a *Access) GetInt8s(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []int8 {
	return []int8{}
}

// GetInt16s ...
func (a *Access) GetInt16s(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []int16 {
	return []int16{}
}

// GetInt32s ...
func (a *Access) GetInt32s(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []int32 {
	return []int32{}
}

// GetInt64s ...
func (a *Access) GetInt64s(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []int64 {
	return []int64{}
}

// GetUints ...
func (a *Access) GetUints(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []uint {
	return []uint{}
}

// GetUint8s ...
func (a *Access) GetUint8s(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []uint8 {
	return []uint8{}
}

// GetUint16s ...
func (a *Access) GetUint16s(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []uint16 {
	return []uint16{}
}

// GetUint32s ...
func (a *Access) GetUint32s(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []uint32 {
	return []uint32{}
}

// GetUint64s ...
func (a *Access) GetUint64s(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []uint64 {
	return []uint64{}
}

// GetBools ...
func (a *Access) GetBools(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []bool {
	return []bool{}
}

// GetTimes ...
func (a *Access) GetTimes(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []time.Time {
	return []time.Time{}
}

// GetStringPtrs ...
func (a *Access) GetStringPtrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*string {
	return []*string{}
}

// GetIntPtrs ...
func (a *Access) GetIntPtrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*int {
	return []*int{}
}

// GetInt8Ptrs ...
func (a *Access) GetInt8Ptrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*int8 {
	return []*int8{}
}

// GetInt16Ptrs ...
func (a *Access) GetInt16Ptrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*int16 {
	return []*int16{}
}

// GetInt32Ptrs ...
func (a *Access) GetInt32Ptrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*int32 {
	return []*int32{}
}

// GetInt64Ptrs ...
func (a *Access) GetInt64Ptrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*int64 {
	return []*int64{}
}

// GetUintPtrs ...
func (a *Access) GetUintPtrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*uint {
	return []*uint{}
}

// GetUint8Ptrs ...
func (a *Access) GetUint8Ptrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*uint8 {
	return []*uint8{}
}

// GetUint16Ptrs ...
func (a *Access) GetUint16Ptrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*uint16 {
	return []*uint16{}
}

// GetUint32Ptrs ...
func (a *Access) GetUint32Ptrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*uint32 {
	return []*uint32{}
}

// GetUint64Ptrs ...
func (a *Access) GetUint64Ptrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*uint64 {
	return []*uint64{}
}

// GetBoolPtrs ...
func (a *Access) GetBoolPtrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*bool {
	return []*bool{}
}

// GetTimePtrs ...
func (a *Access) GetTimePtrs(lock string, filter *jsonapi.Condition, pageSize, pageNumber int) []*time.Time {
	return []*time.Time{}
}

// WillSet ...
func (a *Access) WillSet(lock string) {}

// SetString ...
func (a *Access) SetString(lock string, v string) {}

// SetInt ...
func (a *Access) SetInt(lock string, v int) {}

// SetInt8 ...
func (a *Access) SetInt8(lock string, v int8) {}

// SetInt16 ...
func (a *Access) SetInt16(lock string, v int16) {}

// SetInt32 ...
func (a *Access) SetInt32(lock string, v int32) {}

// SetInt64 ...
func (a *Access) SetInt64(lock string, v int64) {}

// SetUint ...
func (a *Access) SetUint(lock string, v uint) {}

// SetUint8 ...
func (a *Access) SetUint8(lock string, v uint8) {}

// SetUint16 ...
func (a *Access) SetUint16(lock string, v uint16) {}

// SetUint32 ...
func (a *Access) SetUint32(lock string, v uint32) {}

// SetUint64 ...
func (a *Access) SetUint64(lock string, v uint64) {}

// SetBool ...
func (a *Access) SetBool(lock string, v bool) {}

// SetTime ...
func (a *Access) SetTime(lock string, v time.Time) {}

// SetStringPtr ...
func (a *Access) SetStringPtr(lock string, v *string) {}

// SetIntPtr ...
func (a *Access) SetIntPtr(lock string, v *int) {}

// SetInt8Ptr ...
func (a *Access) SetInt8Ptr(lock string, v *int8) {}

// SetInt16Ptr ...
func (a *Access) SetInt16Ptr(lock string, v *int16) {}

// SetInt32Ptr ...
func (a *Access) SetInt32Ptr(lock string, v *int32) {}

// SetInt64Ptr ...
func (a *Access) SetInt64Ptr(lock string, v *int64) {}

// SetUintPtr ...
func (a *Access) SetUintPtr(lock string, v *uint) {}

// SetUint8Ptr ...
func (a *Access) SetUint8Ptr(lock string, v *uint8) {}

// SetUint16Ptr ...
func (a *Access) SetUint16Ptr(lock string, v *uint16) {}

// SetUint32Ptr ...
func (a *Access) SetUint32Ptr(lock string, v *uint32) {}

// SetUint64Ptr ...
func (a *Access) SetUint64Ptr(lock string, v *uint64) {}

// SetBoolPtr ...
func (a *Access) SetBoolPtr(lock string, v *bool) {}

// SetTimePtr ...
func (a *Access) SetTimePtr(lock string, v *time.Time) {}

// SetStrings ...
func (a *Access) SetStrings(lock string, v []string) {}

// SetInts ...
func (a *Access) SetInts(lock string, v []int) {}

// SetInt8s ...
func (a *Access) SetInt8s(lock string, v []int8) {}

// SetInt16s ...
func (a *Access) SetInt16s(lock string, v []int16) {}

// SetInt32s ...
func (a *Access) SetInt32s(lock string, v []int32) {}

// SetInt64s ...
func (a *Access) SetInt64s(lock string, v []int64) {}

// SetUints ...
func (a *Access) SetUints(lock string, v []uint) {}

// SetUint8s ...
func (a *Access) SetUint8s(lock string, v []uint8) {}

// SetUint16s ...
func (a *Access) SetUint16s(lock string, v []uint16) {}

// SetUint32s ...
func (a *Access) SetUint32s(lock string, v []uint32) {}

// SetUint64s ...
func (a *Access) SetUint64s(lock string, v []uint64) {}

// SetBools ...
func (a *Access) SetBools(lock string, v []bool) {}

// SetTimes ...
func (a *Access) SetTimes(lock string, v []time.Time) {}

// SetStringPtrs ...
func (a *Access) SetStringPtrs(lock string, v []*string) {}

// SetIntPtrs ...
func (a *Access) SetIntPtrs(lock string, v []*int) {}

// SetInt8Ptrs ...
func (a *Access) SetInt8Ptrs(lock string, v []*int8) {}

// SetInt16Ptrs ...
func (a *Access) SetInt16Ptrs(lock string, v []*int16) {}

// SetInt32Ptrs ...
func (a *Access) SetInt32Ptrs(lock string, v []*int32) {}

// SetInt64Ptrs ...
func (a *Access) SetInt64Ptrs(lock string, v []*int64) {}

// SetUintPtrs ...
func (a *Access) SetUintPtrs(lock string, v []*uint) {}

// SetUint8Ptrs ...
func (a *Access) SetUint8Ptrs(lock string, v []*uint8) {}

// SetUint16Ptrs ...
func (a *Access) SetUint16Ptrs(lock string, v []*uint16) {}

// SetUint32Ptrs ...
func (a *Access) SetUint32Ptrs(lock string, v []*uint32) {}

// SetUint64Ptrs ...
func (a *Access) SetUint64Ptrs(lock string, v []*uint64) {}

// SetBoolPtrs ...
func (a *Access) SetBoolPtrs(lock string, v []*bool) {}

// SetTimePtrs ...
func (a *Access) SetTimePtrs(lock string, v time.Time) {}

// Release ...
func (a Access) Release(lock string, keep ...[]string) {}

// SetToOneRel ...
func (a *Access) SetToOneRel(key string, id string) {}

// AddToManyRel ...
func (a *Access) AddToManyRel(key string, ids ...string) {}

// SetToManyRel ...
func (a *Access) SetToManyRel(key string, ids []string) {}

// DeleteToManyRel ...
func (a *Access) DeleteToManyRel(key string, id string) {}
