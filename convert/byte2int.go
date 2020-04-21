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
