// Copyright 2020 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package convert
import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//整型转换成字节
func Int2Byte(n int32) []byte {
	//var bytesBuffer = new(bytes.Buffer)
	var bytesBuffer = bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, n)
	if err != nil {
		fmt.Println("Int2Byte func failed:", err)
	}
	return bytesBuffer.Bytes()
}

//字节转换成整型
func Byte2Int(b []byte) (x int32) {
	var byteBuffer = bytes.NewBuffer(b)
	err := binary.Read(byteBuffer, binary.BigEndian, &x)
	if err != nil {
		fmt.Println("Byte2Int func failed:", err)
	}
	return
}
