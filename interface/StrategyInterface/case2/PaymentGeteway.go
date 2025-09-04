//场景：电商网站集成多种支付方式
// 一个电商网站需要支持多种支付方式，比如 Stripe、PayPal、支付宝等。
// 每种支付方式都有自己独特的配置（如 API Key, Secret）和一系列操作（支付、退款、查询状态）。

// 支付方式的策略接口

package main

import (
	"fmt"
	"time"
)

type TransactionStatus string

const (
	StatusSuccess TransactionStatus = "SUCCESS"
	StatusFailed  TransactionStatus = "FAILED"
	StatusPending TransactionStatus = "PENDING"
)

// GeteWay 定义了一组支付行为
// 任何支付行为需要实现下面的三个方法

type PaymentGateway interface {
	Pay(amount float64, userID string) (string, error)
	Refund(amount float64, userID string) error
	GetStatus(transactionID string) (TransactionStatus, error)
}

//现在我们为 Alipay 和 PayPal 创建具体的结构体。
//关键在于，这些结构体需要持有自己的配置信息（API Keys），这就是“状态”。

type Alipay struct {
	APIKey    string
	SecretKey string
}

// 实现支付网关的策略，执行支付行为
func (a *Alipay) Pay(amount float64, userID string) (string, error) {
	fmt.Println("----使用支付宝网关处理支付----")
	fmt.Println("验证API Key:", a.APIKey)
	fmt.Println("验证Secret Key:", a.SecretKey)
	fmt.Printf("对用户 %s 支付 %f 元\n", userID, amount)
	transactionID := fmt.Sprintf("stripe_tx_%d", time.Now().UnixNano())
	fmt.Printf("支付成功，交易ID: %s\n", transactionID)

	// 返回成功状态
	return transactionID, nil
}

// 实现 Refund 方法
func (s *Alipay) Refund(amount float64, userID string) error {
	fmt.Printf("--- 使用 Stripe 网关处理退款 ---\n")
	fmt.Printf("向 Stripe 发起对交易 %s 的 %.2f 元退款...\n", userID, amount)
	return nil
}

// 实现 GetStatus 方法
func (s *Alipay) GetStatus(transactionID string) (TransactionStatus, error) {
	fmt.Printf("--- 使用 Stripe 网关查询状态 ---\n")
	fmt.Printf("向 Stripe 查询交易 %s 的状态...\n", transactionID)
	return StatusSuccess, nil
}

// PayPalGateway 是 PaymentGateway 接口的另一个具体实现。
type PayPalGateway struct {
	ClientID     string
	ClientSecret string
}

// 实现 Pay 方法
func (p *PayPalGateway) Pay(amount float64, userID string) (string, error) {
	fmt.Printf("--- 使用 PayPal 网关处理支付 ---\n")
	fmt.Printf("验证 Client ID: %s\n", p.ClientID)
	fmt.Printf("跳转到 PayPal 对用户 %s 进行 %.2f 元扣款...\n", userID, amount)
	transactionID := fmt.Sprintf("paypal_tx_%d", time.Now().UnixNano())
	fmt.Printf("支付成功，交易ID: %s\n", transactionID)
	return transactionID, nil
}

// 实现 Refund 方法
func (p *PayPalGateway) Refund(amount float64, userID string) error {
	fmt.Printf("--- 使用 PayPal 网关处理退款 ---\n")
	// ...
	return nil
}

// 实现 GetStatus 方法
func (p *PayPalGateway) GetStatus(transactionID string) (TransactionStatus, error) {
	fmt.Printf("--- 使用 PayPal 网关查询状态 ---\n")
	// ...
	return StatusSuccess, nil
}

//------------------------Context------------------------

type CheckoutHandler struct {
	Gateway PaymentGateway
}

func NewCheckoutHandler(gateway PaymentGateway) *CheckoutHandler {
	return &CheckoutHandler{Gateway: gateway}
}

func (h *CheckoutHandler) ProcessPayment(amount float64, userID string) error {
	fmt.Printf("\n>>> 结账处理器开始处理 %.2f 元的订单...\n", amount)
	transactionID, err := h.Gateway.Pay(amount, userID)
	if err != nil {
		fmt.Printf(">>> 支付失败: %v\n", err)
		return err
	}
	fmt.Printf(">>> 支付流程成功，数据库已记录交易ID: %s\n", transactionID)

	return nil
}

func main() {
	// --- 场景一：网站的主要支付方式是 Stripe ---
	fmt.Println("===== 场景一：用户选择 Stripe 支付 =====")
	// 1. 创建一个具体的、有状态的 Stripe 策略实例
	Alipay := &Alipay{
		APIKey:    "alipay_api_key",
		SecretKey: "alipay_secret_key",
	}

	// 2. 创建一个结账处理器，并注入 Stripe 网关
	checkoutWithStripe := NewCheckoutHandler(Alipay)

	// 3. 执行支付流程
	checkoutWithStripe.ProcessPayment(199.99, "user-001")

	// --- 场景二：用户选择了 PayPal 作为支付方式 ---
	fmt.Println("\n\n===== 场景二：用户切换到 PayPal 支付 =====")
	// 1. 创建一个具体的、有状态的 PayPal 策略实例
	payPalGateway := &PayPalGateway{
		ClientID:     "paypal_client_67890",
		ClientSecret: "paypal_secret_fghij",
	}

	// 2. 创建一个新的结账处理器，并注入 PayPal 网关
	// 注意：CheckoutHandler 的代码完全没有变，我们只是换了一个“插件”
	checkoutWithPayPal := NewCheckoutHandler(payPalGateway)

	// 3. 执行完全相同的支付流程
	checkoutWithPayPal.ProcessPayment(49.50, "user-002")
}
