package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/tcp/chatroom/common/message"
)

//将方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //传输时使用的缓冲

}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8094)

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("read head error")
		return
	}
	//fmt.Println("读到的数据是buf=", buf[:n])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}
	//把pkglen反序列化成message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json unmarshal failed err:", err)
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送对方一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("send is failed err:", err)
		return
	}
	//发送数据本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("send is failed err:", err)
		return
	}
	return
}
