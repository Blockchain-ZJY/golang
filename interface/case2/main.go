package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// ===================================================================
// 1. 接口契约 (由 Go 标准库 io 包提供)
// ===================================================================
//
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }
//
// type Writer interface {
//     Write(p []byte) (n int, err error)
// }
//
// 我们将实现自己的 Reader，并使用标准库提供的 Writer。

// ===================================================================
// 2. 自定义一个具体的 Reader 类型
// ===================================================================

// SlowStringReader 是一个 io.Reader 接口的具体实现。
// 它从一个内部字符串中读取数据，但每次只读取一小部分，并模拟延迟。
// 这可以用来模拟从网络连接等慢速数据源读取数据的过程。
type SlowStringReader struct {
	data string // 内部持有的数据源字符串
	pos  int    // 记录当前读取到了哪个位置 (这是内部状态)
}

// 为 *SlowStringReader 类型实现 Read 方法，使其遵守 io.Reader 接口的契约。
// Read 方法会尝试将数据填充到 p 这个字节切片中。
func (r *SlowStringReader) Read(p []byte) (n int, err error) {
	// --- 检查是否已经读完 ---
	// 如果当前位置已经大于或等于字符串的总长度，说明所有数据都已被读取。
	// 我们返回 0 和 io.EOF 信号，这是 Reader 契约中表示“读取结束”的标准方式。
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}

	// --- 模拟耗时操作 ---
	// 模拟每次读取都需要一点时间，比如网络延迟或磁盘寻道。
	time.Sleep(100 * time.Millisecond)

	// --- 执行读取操作 ---
	// Go 内置的 copy 函数非常适合这个场景。
	// 它会从源 (r.data 从当前位置开始的切片) 拷贝数据到目的地 (p)。
	// 它返回实际拷贝的字节数，这个数是源剩余长度和目的地容量中的较小者。
	// 这样可以确保我们不会写入超出 p 容量的数据。
	n = copy(p, r.data[r.pos:])

	// --- 更新内部状态 ---
	// 将当前位置向前移动 n 个字节，为下一次读取做准备。
	r.pos += n

	// 返回读取的字节数和 nil 错误（表示本次读取成功）
	return n, nil
}

// ===================================================================
// 3. 组装和应用
// ===================================================================

func main() {
	// --- 1. 创建具体的“礼物”实例 ---

	// 创建一个我们自定义的 Reader 实例。
	// 它是一个具体的 *SlowStringReader 类型，但它遵守了 io.Reader 接口的契约。
	source := &SlowStringReader{
		data: "你好，这是一个接口的复杂示例...它会一个字一个字地出现。",
	}

	// 使用一个标准库提供的 Writer 实例。
	// os.Stdout 代表程序的标准输出（通常是你的终端），它是一个 *os.File 类型。
	// *os.File 类型已经为我们实现了 io.Writer 接口。
	destination := os.Stdout

	// --- 2. 调用只认接口“盒子”的函数 ---

	fmt.Println("准备开始拷贝数据... 你会看到文字缓慢地出现：")

	// io.Copy 是一个完美的“面向接口编程”的例子。
	// 它不关心 source 是什么，只知道它是一个 Reader。
	// 它不关心 destination 是什么，只知道它是一个 Writer。
	// 它会循环地从 source 读取数据，然后写入 destination，直到 source 返回 io.EOF。
	bytesCopied, err := io.Copy(destination, source)

	// --- 3. 处理结果 ---

	if err != nil {
		// 理论上，我们的实现不会出错，但好的编程习惯是总是检查错误。
		fmt.Printf("\n拷贝过程中发生错误: %v\n", err)
		return
	}

	// 在 io.Copy 结束后，打印一条最终信息。
	// 两个换行符是为了与缓慢输出的文本分开，看起来更清晰。
	fmt.Printf("\n\n拷贝完成！总共拷贝了 %d 字节。\n", bytesCopied)
}
