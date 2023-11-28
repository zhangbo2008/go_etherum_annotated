// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"sync"

	"github.com/holiman/uint256"
) //使用256版本的整数.

// 一个函数把 interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数 返回 interface{},那么也就可以返回任意类型的值。是不是很有用啊！
var stackPool = sync.Pool{ // pool里面设置好new函数就行.
	New: func() interface{} {
		return &Stack{data: make([]uint256.Int, 0, 16)} // 池子里有16个stack,返回类型*stack
	},
}

// Stack is an object for basic stack operations. Items popped to the stack are
// expected to be changed and modified. stack does not take care of adding newly
// initialised objects.
type Stack struct {
	data []uint256.Int
}

func newstack() *Stack {
	return stackPool.Get().(*Stack) //.(*Stack) 是接口转换, 强制转成*Stack指针. //get底层调用stackpool.New.所以返回的是stack的指针. 28行.
}

func returnStack(s *Stack) { //把s放回pool中.
	s.data = s.data[:0] //清空data里面数据.
	stackPool.Put(s)    //放回去.
}

// Data returns the underlying uint256.Int array.
func (st *Stack) Data() []uint256.Int {
	return st.data
}

func (st *Stack) push(d *uint256.Int) {
	// NOTE push limit (1024) is checked in baseCheck
	st.data = append(st.data, *d)
}

func (st *Stack) pop() (ret uint256.Int) {
	ret = st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return
}

func (st *Stack) len() int {
	return len(st.data)
}

func (st *Stack) swap(n int) { //栈顶元素跟距离栈顶n的元素交换.
	st.data[st.len()-n], st.data[st.len()-1] = st.data[st.len()-1], st.data[st.len()-n]
}

func (st *Stack) dup(n int) {
	st.push(&st.data[st.len()-n])
}

func (st *Stack) peek() *uint256.Int { //注意返回的是指针.
	return &st.data[st.len()-1]
}

// Back returns the n'th item in stack
func (st *Stack) Back(n int) *uint256.Int {
	return &st.data[st.len()-n-1]
}
