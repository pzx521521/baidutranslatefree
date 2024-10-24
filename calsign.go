package baidutranslatefree

import (
	"github.com/robertkrimen/otto"
)

// Signer 结构体用于处理签名计算
type Signer struct {
	vm *otto.Otto // Otto 解释器实例
}

// NewSigner 创建一个新的 Signer 实例
func NewSigner() (*Signer, error) {
	vm := otto.New()

	// 执行 JavaScript 代码
	if _, err := vm.Run(JS); err != nil {
		return nil, err
	}

	return &Signer{vm: vm}, nil
}

// CalSign 计算签名
func (s *Signer) CalSign(text string) (string, error) {
	// 调用 JavaScript 函数 a
	value, err := s.vm.Call("token", nil, text)
	if err != nil {
		return "", err
	}

	// 获取返回值
	result, err := value.ToString()
	if err != nil {
		return "", err
	}

	return result, nil
}
