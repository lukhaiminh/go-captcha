/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package v2

import (
	"github.com/lukhaiminh/go-captcha/v2/click"
	"github.com/lukhaiminh/go-captcha/v2/rotate"
	"github.com/lukhaiminh/go-captcha/v2/slide"
)

// NewClickBuilder .
func NewClickBuilder(opts ...click.Option) click.Builder {
	return click.NewBuilder(opts...)
}

// NewSlideBuilder .
func NewSlideBuilder(opts ...slide.Option) slide.Builder {
	return slide.NewBuilder(opts...)
}

// NewRotateBuilder .
func NewRotateBuilder(opts ...rotate.Option) rotate.Builder {
	return rotate.NewBuilder(opts...)
}
