/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (https://www.swig.org).
 * Version 4.1.0
 *
 * Do not make changes to this file unless you know what you are doing - modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

// source: demo.i

package solib

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stddef.h>
#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_solib_a86ad6ccfe6d4225(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_solib_a86ad6ccfe6d4225(swig_intgo arg1);
extern uintptr_t _wrap_new_Demo_solib_a86ad6ccfe6d4225(void);
extern void _wrap_delete_Demo_solib_a86ad6ccfe6d4225(uintptr_t arg1);
extern void _wrap_Demo_SayHello_solib_a86ad6ccfe6d4225(uintptr_t arg1);
extern swig_intgo _wrap_add_solib_a86ad6ccfe6d4225(swig_intgo arg1, swig_intgo arg2);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


func getSwigcptr(v interface { Swigcptr() uintptr }) uintptr {
	if v == nil {
		return 0
	}
	return v.Swigcptr()
}


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_solib_a86ad6ccfe6d4225(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_solib_a86ad6ccfe6d4225(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrDemo uintptr

func (p SwigcptrDemo) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDemo) SwigIsDemo() {
}

func NewDemo() (_swig_ret Demo) {
	var swig_r Demo
	swig_r = (Demo)(SwigcptrDemo(C._wrap_new_Demo_solib_a86ad6ccfe6d4225()))
	return swig_r
}

func DeleteDemo(arg1 Demo) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_Demo_solib_a86ad6ccfe6d4225(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrDemo) SayHello() {
	_swig_i_0 := arg1
	C._wrap_Demo_SayHello_solib_a86ad6ccfe6d4225(C.uintptr_t(_swig_i_0))
}

type Demo interface {
	Swigcptr() uintptr
	SwigIsDemo()
	SayHello()
}

func Add(arg1 int, arg2 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int)(C._wrap_add_solib_a86ad6ccfe6d4225(C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}


