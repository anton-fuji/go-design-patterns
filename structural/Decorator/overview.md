# Decorator パターン
Decoratorパターンは、オブジェクトに対して動的に機能を追加するための構造関連のデザインパターン。
ラッパーオブジェクトを重ねることで機能を拡張できる。

Goではembeddingを活用することで実装が可能

次の点には注意したい。
- 小さなオブジェクトが多数生成される
  - メモリ使用量が増加する可能性
- デバッグの複雑さ
  - ラップの階層が深いとスタックトレースが複雑になる
- 順序依存性
  - デコレータの適用順序によって結果が変わる場合がある

## Beforeコード
**問題点**
- トッピングの組み合わせごとに構造体を作っている
  - トッピングを増やすたびに同じような構造体は増やしたくない
- 拡張性の欠如
  - 新しいトッピングを追加するたびに大量の構造体を作成・修正適用が困難
- 柔軟性の欠如
  - ダブルミルクのような同じトッピングの複数適用が困難

## Afterコード
- 各トッピングで独立したデコレータとして実装している
- デコレータを自由に組み合わせて機能を追加
- 新しいトッピングを追加したい時は、新しいデコレータのファイルを追加するだけ

# アーキテクチャ
```md
┌─────────────┐
│  Beverage   │ (interface)
├─────────────┤
│ Cost()      │
│Description()│
└─────────────┘
       △
       │ implements
       ├──────────────┬────────────────┐
       │              │                │
┌──────┴──────┐ ┌─────┴────────────┐   │
│   Coffee    │ │CondimentDecorator│   │
├─────────────┤ ├──────────────────┤   │
│ name        │ │ beverage         │   │
├─────────────┤ ├──────────────────┤   │
│ Cost()      │ │                  │   │
│Description()│ │                  │   │
└─────────────┘ └──────────────────┘   │
                         △             │
                         │ extends     │
                ┌────────┼───────┬─────┘
                │        │       │
           ┌────┴───┐   ┌┴────┐ ┌┴────┐
           │  Milk  │   │Sugar│ │Whip │
           └────────┘   └─────┘ └─────┘
```


# 実装の流れ
## 1. Beverageインターフェースを作成
```go
type Beverage interface {
    Cost() int
    Description() string
}

```
全ての飲み物が実装すべきインターフェース

## 2. CondimentDecoratorの作成
デコレータの基盤を作成
```go
type CondimentDecorator struct {
    beverage Beverage
}
```


## 3. 具体的なデコレータを実装
今回は、Milk, Sugar, Whipデコレータを用意
- [milk.go](https://github.com/anton-fuji/go-design-patterns/blob/main/structural/Decorator/after/milk.go)
- [sugar.go](https://github.com/anton-fuji/go-design-patterns/blob/main/structural/Decorator/after/sugar.go)
- [whip.go](https://github.com/anton-fuji/go-design-patterns/blob/main/structural/Decorator/after/whip.go)

## 使用例
```go
	// シンプルなコーヒー
	coffee := NewCoffee("Espresso")
	fmt.Printf("%s: %d円\n", coffee.Description(), coffee.Cost())

	// ミルク入り
	coffeeWithMilk := NewMilk(NewCoffee("Espresso"))
	fmt.Printf("%s: %d円\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())

	// ミルクと砂糖入り
	coffeeWithMilkAndSugar := NewSugar(NewMilk(NewCoffee("Espresso")))
	fmt.Printf("%s: %d円\n", coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())

	// ダブルミルク、砂糖、ホイップ入り
	fancyCoffee := NewWhip(
		NewSugar(
			NewMilk(
				NewMilk(
					NewCoffee("Espresso"),
				),
			),
		),
	)
	fmt.Printf("%s: %d円\n", fancyCoffee.Description(), fancyCoffee.Cost())
```


